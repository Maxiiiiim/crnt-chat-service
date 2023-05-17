package chat

import (
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendReply(ctx context.Context, request *pb.SendReplyRequest) (*pb.SendReplyResponse, error) {
	message, err := s.Messages.SentReply(ctx, request.GetSenderId(), request.GetDialogId(), request.GetReplyToId(), request.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if message == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("sender, dialog, user, or message (if any) is not found"))
	}

	message.GetContent().Content = nil
	response := new(pb.SendReplyResponse)
	response.Message = message
	return response, nil
}
