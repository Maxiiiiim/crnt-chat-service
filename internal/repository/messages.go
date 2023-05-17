package repository

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
)

type Messages interface {
	DeleteMessage(ctx context.Context, id string, userID string) error
	GetDialogMessages(ctx context.Context, dialogId string, offsetId string, limit uint64, userID string) ([]*pb.Message, error)
	SendMessage(ctx context.Context, message *pb.Message) (*pb.Message, error)
	UpdateMessage(ctx context.Context, id string, content *pb.MessageContent, userID string) (*pb.Message, error)
	SentReply(ctx context.Context, senderID string, dialogID string, messageID string, content *pb.MessageContent) (*pb.Message, error)
	GetReplies(ctx context.Context, messageID string, offsetID string, limit uint64, userID string) ([]*pb.Message, error)
	SearchMessages(ctx context.Context, dialogID string, query string, offsetID string, limit uint64, userID string) ([]*pb.Message, error)
}
