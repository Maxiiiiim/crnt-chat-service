package chat

import (
	"github.com/Maxiiiiim/crnt-chat-service/internal/repository"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	*Config
}

type Config struct {
	Users    repository.Users
	Dialogs  repository.Dialogs
	Messages repository.Messages
}

func NewServer(config *Config) *Server {
	return &Server{
		Config: config,
	}
}
