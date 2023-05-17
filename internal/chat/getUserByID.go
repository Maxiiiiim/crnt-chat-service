package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserById(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.Users.GetUserByID(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if user == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no user with id %s", request.UserId))
	}
	response := &pb.GetUserByIdResponse{
		User: user,
	}
	return response, nil
}
