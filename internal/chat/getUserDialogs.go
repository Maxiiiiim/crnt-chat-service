package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserDialogs(ctx context.Context, request *pb.GetUserDialogsRequest) (*pb.GetUserDialogsResponse, error) {
	dialogs, err := s.Dialogs.GetUserDialogs(ctx, request.UserId, request.Limit, request.OffsetId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if dialogs == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user or dialog (if any) not fonud"))
	}
	response := new(pb.GetUserDialogsResponse)
	response.Dialogs = dialogs
	return response, nil
}
