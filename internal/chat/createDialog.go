package chat

import (
	"context"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDialog(ctx context.Context, request *pb.CreateDialogRequest) (*pb.CreateDialogResponse, error) {
	dialog, err := s.Dialogs.CreateDialog(ctx, request.GetCreatorId(), request.GetMembers(), request.GetMeta(), request.GetPersonal())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(pb.CreateDialogResponse)
	response.Dialog = dialog
	return response, nil
}
