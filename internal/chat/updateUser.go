package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateUserMeta(ctx context.Context, request *pb.UpdateUserMetaRequest) (*pb.UpdateUserMetaResponse, error) {
	fmt.Println(request.Meta)
	user, err := s.Users.UpdateUserMeta(ctx, request.UserId, request.Meta)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if user == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no user with id %s", request.UserId))
	}
	response := &pb.UpdateUserMetaResponse{
		User: user,
	}
	return response, nil
}
