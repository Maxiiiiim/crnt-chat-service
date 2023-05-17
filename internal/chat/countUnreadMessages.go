package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CountUnreadMessages(ctx context.Context, request *pb.CountUnreadMessagesRequest) (*pb.CountUnreadMessagesResponse, error) {
	count, err := s.Dialogs.CountUnreadMessages(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if count < 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user with id %s is not found", request.UserId))
	}
	response := new(pb.CountUnreadMessagesResponse)
	response.UnreadCount = count
	return response, nil
}
