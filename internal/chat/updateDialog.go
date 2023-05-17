package chat

import (
	"context"
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateDialogMeta(ctx context.Context, request *pb.UpdateDialogMetaRequest) (*pb.UpdateDialogMetaResponse, error) {
	dialog, err := s.Dialogs.UpdateDialogMeta(ctx, request.DialogId, request.Meta)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if dialog == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no dialog with id %s", request.DialogId))
	}
	response := new(pb.UpdateDialogMetaResponse)
	response.Dialog = dialog
	return response, nil
}
