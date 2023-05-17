package chat

import (
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateMessage(ctx context.Context, request *pb.UpdateMessageRequest) (*pb.UpdateMessageResponse, error) {
	message, err := s.Messages.UpdateMessage(ctx, request.MessageId, request.Content, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if message == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("sender, dialog, or message (if any) is not found"))
	}
	response := new(pb.UpdateMessageResponse)
	response.Message = message
	return response, nil
}
