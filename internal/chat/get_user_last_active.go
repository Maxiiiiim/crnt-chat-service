package chat

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetUsersLastActive(ctx context.Context, request *pb.GetUsersLastActiveRequest) (*pb.GetUsersLastActiveResponse, error) {
	m, err := s.Users.GetUserLastActive(ctx, request.GetUserIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.GetUsersLastActiveResponse{
		UsersLastActive: make(map[string]*timestamppb.Timestamp, len(m)),
	}

	for k, v := range m {
		response.UsersLastActive[k] = timestamppb.New(v)
	}

	return response, nil
}
