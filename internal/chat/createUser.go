package chat

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.Users.CreateUser(ctx, request.Meta, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.CreateUserResponse{
		User: user,
	}

	return response, nil
}
