package chat

import (
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReplies(ctx context.Context, request *pb.GetRepliesRequest) (*pb.GetRepliesResponse, error) {
	messages, err := s.Messages.GetReplies(ctx, request.GetMessageId(), request.GetOffsetId(), request.GetLimit(), request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if messages == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no messages with reply_to_id %s", request.MessageId))
	}
	response := new(pb.GetRepliesResponse)
	response.Messages = messages
	return response, nil
}
