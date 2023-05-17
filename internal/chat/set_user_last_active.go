package chat

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SetLastActive(ctx context.Context, request *pb.SetLastActiveRequest) (*pb.SetLastActiveResponse, error) {
	err := s.Users.SetLastActive(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.SetLastActiveResponse{}
	return response, nil
}
