package chat

import (
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteMessage(ctx context.Context, request *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	err := s.Messages.DeleteMessage(ctx, request.MessageId, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(pb.DeleteMessageResponse)
	return response, nil
}
