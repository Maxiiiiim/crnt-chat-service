package chat

import (
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *Server) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	requestMessage := &pb.Message{
		SenderId: request.SenderId,
		DialogId: request.DialogId,
		//ReplyId:        request.ReplyId,
		Content: request.Content,
		SentAt:  timestamppb.New(time.Now()),
	}

	message, err := s.Messages.SendMessage(ctx, requestMessage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if message == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("sender, dialog, user, or message (if any) is not found"))
	}

	message.GetContent().Content = nil
	response := new(pb.SendMessageResponse)
	response.Message = message
	return response, nil
}
