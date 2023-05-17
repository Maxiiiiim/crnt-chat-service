package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LeaveDialog(ctx context.Context, request *pb.LeaveDialogRequest) (*pb.LeaveDialogResponse, error) {
	dialog, err := s.Dialogs.LeaveDialog(ctx, request.DialogId, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if dialog == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("dialog or user not found"))
	}
	response := new(pb.LeaveDialogResponse)
	response.Dialog = dialog
	return response, nil
}
