package dialogs

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (r *Repository) CreateDialog(ctx context.Context, creatorID string, members []string, meta map[string]string, personal bool) (*pb.Dialog, error) {
	dialogID := uuid.New().String()
	messageID := uuid.New().String()
	found := false
	for _, member := range members {
		if member == creatorID {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("creator is not a member of this dialog")
	}

	if personal && len(members) != 2 {
		return nil, errors.New("personal dialog must have exactly 2 members")
	}

	if meta == nil {
		meta = make(map[string]string)
	}

	now := time.Now()
	message := &pb.Message{
		Id:       messageID,
		DialogId: dialogID,
		Content: &pb.MessageContent{
			Content: &pb.MessageContent_ServiceContent{
				ServiceContent: &pb.ServiceContent{},
			},
			Text: "Чат создан",
		},
	}

	dialog := &pb.Dialog{
		Id:       dialogID,
		Members:  members,
		Meta:     meta,
		Personal: personal,
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}

	_, err = tx.Exec(ctx, `INSERT INTO dialogs (id, meta, last_message_id, personal) VALUES ($1, $2, $3, $4)`,
		dialog.Id, dialog.Meta, message.Id, personal,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't insert dialog: %w", err)
	}

	_, err = tx.Exec(ctx, `INSERT INTO messages (id, dialog_id, sent_at, updated_at, content_type, content_text)
		VALUES ($1, $2, $3, $3, $4, $5)`,
		message.Id, message.DialogId, now, 4, message.Content.Text,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't insert message: %w", err)
	}

	for _, memberId := range members {
		_, err = tx.Exec(ctx, `INSERT INTO dialog_members (user_id, dialog_id, is_owner) VALUES ($1, $2, $3)`,
			memberId, dialog.Id, memberId == creatorID,
		)
		if err != nil {
			_ = tx.Rollback(ctx)
			return nil, fmt.Errorf("can't insert dialog_members: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't commit transaction: %w", err)
	}

	return dialog, nil
}

func (r *Repository) GetDialogByID(ctx context.Context, dialogID string, userID string) (*pb.Dialog, error) {
	_, err := uuid.Parse(dialogID)
	if err != nil {
		return nil, fmt.Errorf("can't parse dialogID: %w", err)
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("can't parse user id: %w", err)
	}

	dialog := &struct {
		*pb.Dialog
		CreatedAt time.Time
	}{
		Dialog: &pb.Dialog{
			LastMessage: &pb.Message{},
		},
	}

	members := make([]*pb.User, 0)

	err = pgxscan.Get(ctx, r.db, dialog, `SELECT id, meta, created_at, personal FROM dialogs WHERE id = $1`, dialogID)
	if err != nil {
		return nil, fmt.Errorf("can't get dialog: %w", err)
	}

	err = pgxscan.Select(ctx, r.db, &members, `SELECT users.id, users.meta FROM users
		JOIN dialog_members dm on users.id = dm.user_id
		WHERE dm.dialog_id = $1`, dialogID,
	)
	if err != nil {
		return nil, fmt.Errorf("can't get dialog members: %w", err)
	}

	found := false
	for _, m := range members {
		dialog.Members = append(dialog.Members, m.Id)
		if m.Id == userID {
			found = true
		}
	}
	if !found {
		return nil, errors.New("user is not a member of this dialog")
	}

	// get last read message
	err = pgxscan.Get(ctx, r.db, &dialog.LastReadMessageId,
		`SELECT last_read_message_id FROM dialog_members JOIN messages m on m.id = last_read_message_id WHERE dialog_members.user_id = $1 AND dialog_members.dialog_id = $2`, userID, dialogID,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("can't get last read message id: %w", err)
	}

	// get unread messages count
	if dialog.LastReadMessageId == "" {
		err = pgxscan.Get(ctx, r.db, &dialog.UnreadCount, `SELECT COUNT(*) FROM messages WHERE dialog_id = $1 AND is_deleted = false AND reply_to_id IS NULL `, dialogID)
	} else {
		err = pgxscan.Get(ctx, r.db, &dialog.UnreadCount, `
			SELECT COUNT(*) FROM messages
			WHERE dialog_id = $1 AND is_deleted = false AND reply_to_id IS NULL AND (sent_at > (select sent_at from messages where id = $2))`, dialogID, dialog.LastReadMessageId)
	}
	if err != nil {
		return nil, fmt.Errorf("can't get unread messages count: %w", err)
	}

	// get last message in dialog
	err = pgxscan.Get(ctx, r.db, &dialog.LastMessage.Id,
		`SELECT id FROM messages WHERE dialog_id = $1 AND is_deleted = false AND reply_to_id IS NULL ORDER BY sent_at DESC LIMIT 1`, dialogID)

	err = r.fillMessageByID(ctx, dialog.LastMessage)
	if err != nil {
		return nil, fmt.Errorf("can't get last message: %w", err)
	}

	dialog.Dialog.CreatedAt = timestamppb.New(dialog.CreatedAt)
	return dialog.Dialog, nil
}

func (r *Repository) fillMessageByID(ctx context.Context, msg *pb.Message) error {
	var message struct {
		ID                string
		SenderID          string
		SentAt            time.Time
		UpdatedAt         time.Time
		ContentType       int
		ContentText       string
		ContentAdditional [][]byte
		ContentMeta       map[string]string
	}

	err := pgxscan.Get(ctx, r.db, &message, `
		SELECT m.id,
			   m.sender_id,
			   m.sent_at,
			   m.updated_at,
			   m.content_type,
			   m.content_text,
			   m.content_additional,
			   m.content_meta
		FROM messages m
		where m.id = $1;`, msg.Id)
	if err != nil {
		return fmt.Errorf("pgxscan.Select: %w", err)
	}

	msg.SenderId = message.SenderID
	msg.SentAt = timestamppb.New(message.SentAt)
	msg.UpdatedAt = timestamppb.New(message.UpdatedAt)
	msg.Content = &pb.MessageContent{
		Text: message.ContentText,
		Meta: message.ContentMeta,
	}

	switch message.ContentType {
	case 1:
		msg.Content.Content = &pb.MessageContent_TextContent{}
	case 2:
		msg.Content.Content = &pb.MessageContent_MediaContent{
			MediaContent: &pb.MediaContent{
				//Media: message.ContentAdditional,
			},
		}
	case 3:
		msg.Content.Content = &pb.MessageContent_FileContent{
			FileContent: &pb.FileContent{
				//Files: message.ContentAdditional,
			},
		}
	case 4:
		msg.Content.Content = &pb.MessageContent_ServiceContent{}
	}

	return nil
}

func (r *Repository) GetUserDialogs(ctx context.Context, userId string, limit int64, offsetId string) ([]*pb.Dialog, error) {
	_, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("can't parse user id: %w", err)
	}

	// get dialog ids sorted by last message timestamp
	dialogIds := make([]string, 0)
	offset := offsetId != ""
	query := `
		SELECT dialog_members.dialog_id
		FROM dialog_members
		JOIN dialogs d on dialog_members.dialog_id = d.id
		JOIN messages m ON m.id = d.last_message_id
		WHERE dialog_members.user_id = $1 `
	subquery := ` AND sent_at < (SELECT sent_at FROM messages WHERE id = (select last_message_id from dialogs where id = $3))`
	end := `ORDER BY m.sent_at DESC LIMIT $2;`
	if offset {
		err = pgxscan.Select(ctx, r.db, &dialogIds, query+subquery+end, userId, limit, offsetId)
	} else {
		err = pgxscan.Select(ctx, r.db, &dialogIds, query+end, userId, limit)
	}
	if err != nil {
		return nil, fmt.Errorf("can't get dialog ids: %w", err)
	}

	dialogs := make([]*pb.Dialog, 0)

	for _, id := range dialogIds {
		dialog, err := r.GetDialogByID(ctx, id, userId)
		if err != nil {
			return nil, fmt.Errorf("can't get dialog: %w", err)
		}

		dialogs = append(dialogs, dialog)
	}

	return dialogs, nil
}

func (r *Repository) JoinDialog(ctx context.Context, dialogId string, userId string) (*pb.Dialog, error) {
	_, err := uuid.Parse(dialogId)
	if err != nil {
		return nil, fmt.Errorf("can't parse dialog id: %w", err)
	}
	_, err = uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("can't parse user id: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}

	dialog := &struct {
		*pb.Dialog
		CreatedAt time.Time
	}{
		Dialog: &pb.Dialog{},
	}
	err = pgxscan.Get(ctx, tx, dialog, `SELECT id, meta, created_at, personal FROM dialogs WHERE id = $1`, dialogId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't get dialog: %w", err)
	}

	if dialog.Personal {
		return nil, fmt.Errorf("can't join personal dialog")
	}

	// check if user exists
	var exists bool
	err = pgxscan.Get(ctx, tx, &exists, `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, userId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't check if user exists: %w", err)
	}
	if !exists {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("user with id %s doesn't exist", userId)
	}

	// check if user is already in dialog
	err = pgxscan.Get(ctx, tx, &exists, `SELECT EXISTS(SELECT 1 FROM dialog_members WHERE dialog_id = $1 AND user_id = $2)`, dialogId, userId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't check if user is already in dialog: %w", err)
	}
	if exists {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("user with id %s is already in dialog with id %s", userId, dialogId)
	}

	// add user to dialog
	err = pgxscan.Get(ctx, tx, &dialog.LastReadMessageId, `
			INSERT INTO dialog_members (dialog_id, user_id, last_read_message_id)
			VALUES ($1, $2, (SELECT id
							 FROM messages
							 WHERE dialog_id = $1
							   AND is_deleted = FALSE
							   AND reply_to_id IS NULL
							 ORDER BY sent_at DESC
							 LIMIT 1))
			RETURNING last_read_message_id
		`, dialogId, userId,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't add user to dialog: %w", err)
	}

	dialog.LastMessage = &pb.Message{
		Id: dialog.LastReadMessageId,
	}

	// get last message in dialog
	err = r.fillMessageByID(ctx, dialog.LastMessage)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't get last message: %w", err)
	}

	// get dialog members
	dialog.Members = make([]string, 0)
	err = pgxscan.Select(ctx, r.db, &dialog.Members, `
		SELECT user_id
		FROM dialog_members
		WHERE dialog_id = $1
		ORDER BY is_owner;`, dialogId,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't get dialog members: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't commit transaction: %w", err)
	}

	dialog.Dialog.CreatedAt = timestamppb.New(dialog.CreatedAt)
	return dialog.Dialog, nil
}

func (r *Repository) LeaveDialog(ctx context.Context, dialogId string, userId string) (*pb.Dialog, error) {
	_, err := uuid.Parse(dialogId)
	if err != nil {
		return nil, fmt.Errorf("can't parse dialog id: %w", err)
	}
	_, err = uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("can't parse user id: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}

	dialog := &struct {
		*pb.Dialog
		CreatedAt time.Time
	}{
		Dialog: &pb.Dialog{},
	}
	err = pgxscan.Get(ctx, tx, dialog, `SELECT id, meta, created_at, personal FROM dialogs WHERE id = $1`, dialogId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't get dialog: %w", err)
	}

	// check if user is already in dialog
	var exists bool
	err = pgxscan.Get(ctx, tx, &exists, `SELECT EXISTS(SELECT 1 FROM dialog_members WHERE dialog_id = $1 AND user_id = $2)`, dialogId, userId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't check if user is already in dialog: %w", err)
	}
	if !exists {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("user with id %s is not in dialog with id %s", userId, dialogId)
	}

	// remove user from dialog
	_, err = tx.Exec(ctx, `DELETE FROM dialog_members WHERE dialog_id = $1 AND user_id = $2`, dialogId, userId)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't remove user from dialog: %w", err)
	}

	dialog.Dialog.CreatedAt = timestamppb.New(dialog.CreatedAt)
	return dialog.Dialog, nil
}

func (r *Repository) UpdateDialogMeta(ctx context.Context, id string, meta map[string]string) (*pb.Dialog, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("can't parse dialog id: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}

	dialog := &pb.Dialog{}
	err = pgxscan.Get(ctx, tx, dialog, `SELECT id, meta FROM dialogs WHERE id = $1`, id)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't get dialog: %w", err)
	}

	// update dialog meta
	_, err = tx.Exec(ctx, `UPDATE dialogs SET meta = $1 WHERE id = $2`, meta, id)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("can't update dialog meta: %w", err)
	}

	dialog.Meta = meta

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't commit transaction: %w", err)
	}

	return dialog, nil
}

func (r *Repository) CountUnreadMessages(ctx context.Context, userId string) (int64, error) {
	_, err := uuid.Parse(userId)
	if err != nil {
		return 0, fmt.Errorf("can't parse user id: %w", err)
	}

	// get count of unread messages from users dialogs
	var count int64
	err = pgxscan.Get(ctx, r.db, &count, `
		WITH x AS (SELECT dialog_members.dialog_id,
			   last_m.sent_at                                          AS last_sent_at,
			   COALESCE(last_seen_m.sent_at, '1970-01-01 00:00:00+00') AS last_seen_sent_at
		FROM dialog_members
				 JOIN dialogs d ON dialog_members.dialog_id = d.id
				 JOIN messages last_m ON last_m.id = d.last_message_id
				 LEFT JOIN messages last_seen_m ON last_seen_m.id = dialog_members.last_read_message_id
		WHERE user_id = $1)
		SELECT COUNT(1)
		FROM x
		WHERE x.LAST_SENT_AT > x.LAST_SEEN_SENT_AT;`, userId)

	return count, nil
}
