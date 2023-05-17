package messages

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (r *Repository) DeleteMessage(ctx context.Context, id string, userID string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid message id")
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	msg := struct {
		ReplyToID sql.NullString
		DialogID  string
		SenderID  string
	}{}

	err = pgxscan.Get(ctx, tx, &msg, `SELECT reply_to_id, dialog_id, sender_id FROM messages WHERE id = $1 and sender_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("message not found: pgxscan.Get: %w", err)
	}

	_, err = tx.Exec(ctx, `UPDATE messages SET is_deleted = true WHERE id = $1 AND sender_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("tx.Exec: %w", err)
	}

	_, err = tx.Exec(ctx, `UPDATE dialogs SET last_message_id = (SELECT id FROM messages WHERE dialog_id = $1 and is_deleted = false AND reply_to_id IS NULL ORDER BY sent_at DESC LIMIT 1) WHERE id = $1`, msg.DialogID)
	if err != nil {
		return fmt.Errorf("tx.Exec: %w", err)
	}

	if msg.ReplyToID.Valid {
		_, err = tx.Exec(ctx, `UPDATE messages SET replies_count = replies_count - 1 WHERE id = $1`, msg.ReplyToID)
		if err != nil {
			return fmt.Errorf("tx.Exec: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}

func (r *Repository) GetDialogMessages(ctx context.Context, dialogId string, offsetId string, limit uint64, userID string) ([]*pb.Message, error) {
	_, err := uuid.Parse(dialogId)
	if err != nil {
		return nil, fmt.Errorf("invalid dialog id")
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}
	_, err = uuid.Parse(offsetId)
	if offsetId != "" && err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	var messages []*struct {
		ID                string
		SenderID          string
		SentAt            time.Time
		UpdatedAt         time.Time
		ContentType       int
		ContentText       string
		ContentAdditional [][]byte
		ContentMeta       map[string]string
		RepliesCount      uint64
		IsUnseenByMe      bool
		IsSeen            bool
	}

	if offsetId == "" {
		offsetId = uuid.Nil.String()
	}

	err = pgxscan.Select(ctx, tx, &messages, `
		SELECT m.id,
			   m.sender_id,
			   m.sent_at,
			   m.updated_at,
			   m.content_type,
			   m.content_text,
			   m.content_additional,
			   m.content_meta,
			   m.replies_count,
			   m.sent_at <= (select coalesce(max(m.sent_at), '1970-01-01 00:00:00+00') from dialog_members dm left join messages m on m.id = dm.last_read_message_id where dm.user_id != $4 limit 1) as is_seen
		FROM messages m
		WHERE dialog_id = $1
		  AND is_deleted = FALSE
		  AND reply_to_id IS NULL
		  AND sent_at < COALESCE((SELECT sent_at FROM messages WHERE id = $2 AND dialog_id = $1), '2999-01-01 00:00:00+00')
		ORDER BY sent_at desc
		LIMIT $3;`, dialogId, offsetId, limit, userID)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	if len(messages) > 0 {
		var lastRead struct {
			LastReadMessageId string
			LastSeenTime      sql.NullTime
		}
		err = pgxscan.Get(ctx, tx, &lastRead, `
		SELECT last_read_message_id,
			   m.sent_at as last_seen_time
		FROM dialog_members
		LEFT JOIN messages m ON dialog_members.last_read_message_id = m.id
		WHERE dialog_members.dialog_id = $1
		  AND user_id = $2;`, dialogId, userID)
		if err != nil {
			return nil, fmt.Errorf("pgxscan.Get: %w", err)
		}

		if lastRead.LastReadMessageId == uuid.Nil.String() || lastRead.LastSeenTime.Time.Before(messages[0].SentAt) {
			_, err = tx.Exec(ctx, `
			UPDATE dialog_members
			SET last_read_message_id = $1
			WHERE dialog_id = $2
			  AND user_id = $3;`, messages[0].ID, dialogId, userID)
			if err != nil {
				return nil, fmt.Errorf("tx.Exec: %w", err)
			}
		}

		for _, m := range messages {
			if lastRead.LastSeenTime.Time.Before(m.SentAt) {
				m.IsUnseenByMe = true
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	var result []*pb.Message
	for _, m := range messages {
		msg := &pb.Message{
			Id:        m.ID,
			DialogId:  dialogId,
			SenderId:  m.SenderID,
			SentAt:    timestamppb.New(m.SentAt),
			UpdatedAt: timestamppb.New(m.UpdatedAt),
			Content: &pb.MessageContent{
				Text: m.ContentText,
				Meta: m.ContentMeta,
			},
			RepliesCount: m.RepliesCount,
			Seen:         m.IsSeen,
			SeenByMe:     !m.IsUnseenByMe,
		}
		switch m.ContentType {
		case 1:
			msg.Content.Content = &pb.MessageContent_TextContent{}
		case 2:
			msg.Content.Content = &pb.MessageContent_MediaContent{
				MediaContent: &pb.MediaContent{
					Media: m.ContentAdditional,
				},
			}
		case 3:
			msg.Content.Content = &pb.MessageContent_FileContent{
				FileContent: &pb.FileContent{
					Files: m.ContentAdditional,
				},
			}
		case 4:
			msg.Content.Content = &pb.MessageContent_ServiceContent{}
		}

		result = append(result, msg)
	}

	return result, nil
}

func (r *Repository) SendMessage(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	id := uuid.New().String()
	if _, err := uuid.Parse(message.SenderId); err != nil {
		return nil, fmt.Errorf("invalid sender id")
	}
	if _, err := uuid.Parse(message.DialogId); err != nil {
		return nil, fmt.Errorf("invalid dialog id")
	}

	contentType := 0
	contentAdditional := [][]byte{}
	switch c := message.GetContent().GetContent().(type) {
	case *pb.MessageContent_TextContent:
		contentType = 1
	case *pb.MessageContent_MediaContent:
		contentType = 2
		contentAdditional = c.MediaContent.GetMedia()
	case *pb.MessageContent_FileContent:
		contentType = 3
		contentAdditional = c.FileContent.GetFiles()
	case *pb.MessageContent_ServiceContent:
		contentType = 4
	}
	if contentType == 0 {
		return nil, fmt.Errorf("invalid content type")
	}

	contentMeta := message.GetContent().GetMeta()
	if contentMeta == nil {
		contentMeta = map[string]string{}
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	var _id string
	err = pgxscan.Get(ctx, tx, &_id, `SELECT dialog_id FROM dialog_members WHERE dialog_id = $1 AND user_id = $2`, message.GetDialogId(), message.GetSenderId())
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user %v is not a member of the dialog %v", message.GetSenderId(), message.GetDialogId())
	} else if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %w", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO messages (id, dialog_id, sender_id, sent_at, updated_at, content_type, content_text, content_additional, content_meta)
		VALUES ($1, $2, $3, $4, $4, $5, $6, $7, $8)
	`, id, message.GetDialogId(), message.GetSenderId(), message.GetSentAt().AsTime(), contentType, message.GetContent().GetText(), contentAdditional, contentMeta)
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: %w", err)
	}

	_, err = tx.Exec(ctx, `
		UPDATE dialogs
		SET last_message_id = $1
		WHERE id = $2`, id, message.GetDialogId())
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	message.Id = id
	return message, nil
}

func (r *Repository) UpdateMessage(ctx context.Context, id string, content *pb.MessageContent, userID string) (*pb.Message, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	contentType := 0
	var contentAdditional = [][]byte{}
	switch c := content.GetContent().(type) {
	case *pb.MessageContent_TextContent:
		contentType = 1
	case *pb.MessageContent_MediaContent:
		contentType = 2
		contentAdditional = c.MediaContent.GetMedia()
	case *pb.MessageContent_FileContent:
		contentType = 3
		contentAdditional = c.FileContent.GetFiles()
	case *pb.MessageContent_ServiceContent:
		contentType = 4
	}
	if contentType == 0 {
		return nil, fmt.Errorf("invalid content type")
	}

	contentMeta := content.GetMeta()
	if contentMeta == nil {
		contentMeta = map[string]string{}
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	message := struct {
		SenderId    string `db:"sender_id"`
		ContentType int    `db:"content_type"`
	}{}
	err = pgxscan.Get(ctx, tx, &message, `SELECT sender_id, content_type FROM messages WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %w", err)
	}

	if message.SenderId != userID {
		return nil, fmt.Errorf("user is not the sender")
	}
	if message.ContentType != contentType {
		return nil, fmt.Errorf("content type mismatch")
	}

	_, err = tx.Exec(ctx, `
		UPDATE messages
		SET content_text = $1, content_additional = $2, content_meta = $3, updated_at = $4
		WHERE id = $5
	`, content.GetText(), contentAdditional, contentMeta, time.Now(), id)
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	return &pb.Message{
		Id:       id,
		SenderId: userID,
	}, nil
}

func (r *Repository) SentReply(ctx context.Context, senderID string, dialogID string, messageID string, content *pb.MessageContent) (*pb.Message, error) {
	_, err := uuid.Parse(senderID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = uuid.Parse(dialogID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = uuid.Parse(messageID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	contentType := 0
	var contentAdditional = [][]byte{}
	switch c := content.GetContent().(type) {
	case *pb.MessageContent_TextContent:
		contentType = 1
	case *pb.MessageContent_MediaContent:
		contentType = 2
		contentAdditional = c.MediaContent.GetMedia()
	case *pb.MessageContent_FileContent:
		contentType = 3
		contentAdditional = c.FileContent.GetFiles()
	case *pb.MessageContent_ServiceContent:
		contentType = 4
	}
	if contentType == 0 {
		return nil, fmt.Errorf("invalid content type")
	}

	contentMeta := content.GetMeta()
	if contentMeta == nil {
		contentMeta = map[string]string{}
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	var _id string
	err = pgxscan.Get(ctx, tx, &_id, `SELECT dialog_id FROM dialog_members WHERE dialog_id = $1 AND user_id = $2`, dialogID, senderID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user %v is not a member of the dialog %v", senderID, dialogID)
	} else if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %w", err)
	}

	message := struct {
		DialogId string         `db:"dialog_id"`
		RelyToId sql.NullString `db:"reply_to_id"`
	}{}
	err = pgxscan.Get(ctx, tx, &message, `SELECT dialog_id, reply_to_id FROM messages WHERE id = $1`, messageID)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %w", err)
	}

	if message.DialogId != dialogID {
		return nil, fmt.Errorf("message is not in dialog")
	}
	if message.RelyToId.Valid {
		return nil, fmt.Errorf("cannot reply to reply")
	}

	id := uuid.New().String()

	_, err = tx.Exec(ctx, `
		INSERT INTO messages (id, dialog_id, sender_id, sent_at, updated_at, content_type, content_text, content_additional, content_meta, reply_to_id)
		VALUES ($1, $2, $3, $4, $4, $5, $6, $7, $8, $9)
	`, id, dialogID, senderID, time.Now(), contentType, content.GetText(), contentAdditional, contentMeta, messageID)
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: %w", err)
	}

	_, err = tx.Exec(ctx, `
		UPDATE messages
		SET replies_count = replies_count + 1
		WHERE id = $1
	`, messageID)
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	return &pb.Message{
		Id:        id,
		DialogId:  dialogID,
		SenderId:  senderID,
		Content:   content,
		ReplyToId: messageID,
	}, nil
}

func (r *Repository) GetReplies(ctx context.Context, messageID string, offsetID string, limit uint64, userID string) ([]*pb.Message, error) {
	_, err := uuid.Parse(messageID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = uuid.Parse(offsetID)
	if err != nil && offsetID != "" {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	var replies []struct {
		ID                string
		DialogID          string
		SenderID          string
		SentAt            time.Time
		UpdatedAt         time.Time
		ContentType       int
		ContentText       string
		ContentAdditional [][]byte
		ContentMeta       map[string]string
		ReplyToID         string
	}

	if offsetID == "" {
		offsetID = uuid.Nil.String()
	}

	err = pgxscan.Select(ctx, tx, &replies, `
		SELECT id, dialog_id, sender_id, sent_at, updated_at, content_type, content_text, content_additional, content_meta, reply_to_id
		FROM messages
		WHERE reply_to_id = $1 
			AND sent_at < COALESCE((SELECT sent_at FROM messages WHERE id = $2), '2999-01-01 00:00:00+00')
			AND is_deleted = false
		ORDER BY sent_at DESC
		LIMIT $3
	`, messageID, offsetID, limit)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	var result []*pb.Message
	for _, m := range replies {
		msg := &pb.Message{
			Id:        m.ID,
			DialogId:  m.DialogID,
			SenderId:  m.SenderID,
			SentAt:    timestamppb.New(m.SentAt),
			UpdatedAt: timestamppb.New(m.UpdatedAt),
			Content: &pb.MessageContent{
				Text: m.ContentText,
				Meta: m.ContentMeta,
			},
			ReplyToId: m.ReplyToID,
		}
		switch m.ContentType {
		case 1:
			msg.Content.Content = &pb.MessageContent_TextContent{}
		case 2:
			msg.Content.Content = &pb.MessageContent_MediaContent{
				MediaContent: &pb.MediaContent{
					Media: m.ContentAdditional,
				},
			}
		case 3:
			msg.Content.Content = &pb.MessageContent_FileContent{
				FileContent: &pb.FileContent{
					Files: m.ContentAdditional,
				},
			}
		case 4:
			msg.Content.Content = &pb.MessageContent_ServiceContent{}
		}

		result = append(result, msg)
	}

	return result, nil
}

func (r *Repository) SearchMessages(ctx context.Context, dialogID string, query string, offsetID string, limit uint64, userID string) ([]*pb.Message, error) {
	_, err := uuid.Parse(dialogID)
	if err != nil {
		return nil, fmt.Errorf("invalid dialog id")
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}
	_, err = uuid.Parse(offsetID)
	if offsetID != "" && err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	var messages []struct {
		ID                string
		DialogID          string
		SenderID          string
		SentAt            time.Time
		UpdatedAt         time.Time
		ContentType       int
		ContentText       string
		ContentAdditional [][]byte
		ContentMeta       map[string]string
		ReplyToID         sql.NullString
		RepliesCount      uint64
	}

	if offsetID == "" {
		offsetID = uuid.Nil.String()
	}

	query = "%" + query + "%"

	err = pgxscan.Select(ctx, tx, &messages, `
		SELECT m.id,
		       m.dialog_id,
			   m.sender_id,
			   m.sent_at,
			   m.updated_at,
			   m.content_type,
			   m.content_text,
			   m.content_additional,
			   m.content_meta,
			   m.reply_to_id,
			   m.replies_count
		FROM messages m
		WHERE dialog_id = $1
		  AND is_deleted = FALSE
		  AND sent_at < COALESCE((SELECT sent_at FROM messages WHERE id = $2 AND dialog_id = $1), '2999-01-01 00:00:00+00')
		  AND (content_text ILIKE $4)
		ORDER BY sent_at desc
		LIMIT $3;`, dialogID, offsetID, limit, query)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit: %w", err)
	}

	var result []*pb.Message
	for _, m := range messages {
		msg := &pb.Message{
			Id:        m.ID,
			DialogId:  dialogID,
			SenderId:  m.SenderID,
			SentAt:    timestamppb.New(m.SentAt),
			UpdatedAt: timestamppb.New(m.UpdatedAt),
			Content: &pb.MessageContent{
				Text: m.ContentText,
				Meta: m.ContentMeta,
			},
			ReplyToId:    m.ReplyToID.String,
			RepliesCount: m.RepliesCount,
		}
		switch m.ContentType {
		case 1:
			msg.Content.Content = &pb.MessageContent_TextContent{}
		case 2:
			msg.Content.Content = &pb.MessageContent_MediaContent{
				MediaContent: &pb.MediaContent{
					Media: m.ContentAdditional,
				},
			}
		case 3:
			msg.Content.Content = &pb.MessageContent_FileContent{
				FileContent: &pb.FileContent{
					Files: m.ContentAdditional,
				},
			}
		case 4:
			msg.Content.Content = &pb.MessageContent_ServiceContent{}
		}

		result = append(result, msg)
	}

	return result, nil
}
