package chat

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"log"
)

func (s *Server) Ping(_ context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	response := new(pb.PingResponse)
	log.Println("ping")
	return response, nil
}
