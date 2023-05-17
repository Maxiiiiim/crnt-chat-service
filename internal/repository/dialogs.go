package repository

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
)

type Dialogs interface {
	CreateDialog(ctx context.Context, creatorID string, members []string, meta map[string]string, personal bool) (*pb.Dialog, error)
	GetDialogByID(ctx context.Context, id string, userID string) (*pb.Dialog, error)
	GetUserDialogs(ctx context.Context, userId string, limit int64, offsetId string) ([]*pb.Dialog, error)
	JoinDialog(ctx context.Context, dialogId string, userId string) (*pb.Dialog, error)
	LeaveDialog(ctx context.Context, dialogId string, userId string) (*pb.Dialog, error)
	UpdateDialogMeta(ctx context.Context, id string, meta map[string]string) (*pb.Dialog, error)
	CountUnreadMessages(ctx context.Context, userId string) (int64, error)
}
