package repository

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"time"
)

type Users interface {
	CreateUser(ctx context.Context, meta map[string]string, id string) (*pb.User, error)
	GetUserByID(ctx context.Context, id string) (*pb.User, error)
	UpdateUserMeta(ctx context.Context, id string, meta map[string]string) (*pb.User, error)
	SetLastActive(ctx context.Context, userID string) error
	GetUserLastActive(ctx context.Context, userIDs []string) (map[string]time.Time, error)
}
