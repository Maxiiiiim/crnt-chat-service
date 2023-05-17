package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDialogById(ctx context.Context, request *pb.GetDialogByIdRequest) (*pb.GetDialogByIdResponse, error) {
	dialog, err := s.Dialogs.GetDialogByID(ctx, request.GetId(), request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if dialog == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no dialog with id %s", request.Id))
	}
	response := new(pb.GetDialogByIdResponse)
	response.Dialog = dialog
	return response, nil
}
