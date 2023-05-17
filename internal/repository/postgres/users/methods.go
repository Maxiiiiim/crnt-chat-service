package users

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, meta map[string]string, id string) (*pb.User, error) {
	if id == "" {
		id = uuid.New().String()
	}
	if len(meta) == 0 {
		meta = make(map[string]string)
	}
	user := &pb.User{
		Id:   id,
		Meta: meta,
	}
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %v", err)
	}
	_, err = tx.Exec(ctx, "INSERT INTO users (id, meta) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET meta = $2", user.Id, user.Meta)
	if err != nil {
		_ = tx.Rollback(ctx)
		return nil, fmt.Errorf("could not insert user: %v", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not commit transaction: %v", err)
	}
	fmt.Println(user)
	return user, err
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*pb.User, error) {
	var user pb.User
	err := r.db.QueryRow(ctx, "SELECT id, meta FROM users WHERE id = $1", id).Scan(&user.Id, &user.Meta)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %v", err)
	}
	return &user, nil
}

func (r *Repository) UpdateUserMeta(ctx context.Context, id string, meta map[string]string) (*pb.User, error) {
	var user pb.User
	err := r.db.QueryRow(ctx, "UPDATE users SET meta = $1 WHERE id = $2 RETURNING id, meta", meta, id).Scan(&user.Id, &user.Meta)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %v", err)
	}
	return &user, nil
}

func (r *Repository) SetLastActive(ctx context.Context, userID string) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET last_active = NOW() WHERE id = $1", userID)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}
	return nil
}

func (r *Repository) GetUserLastActive(ctx context.Context, userIDs []string) (map[string]time.Time, error) {
	for _, id := range userIDs {
		_, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("could not parse uuid %s: %w", id, err)
		}
	}
	rows, err := r.db.Query(ctx, "SELECT id, last_active FROM users WHERE id = ANY($1)", userIDs)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %v", err)
	}
	defer rows.Close()
	users := make(map[string]time.Time)
	for rows.Next() {
		var userID string
		var lastActive time.Time
		err := rows.Scan(&userID, &lastActive)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		users[userID] = lastActive
	}
	return users, nil

}
