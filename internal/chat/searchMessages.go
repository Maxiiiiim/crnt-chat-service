package chat

import (
	"fmt"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchMessages(ctx context.Context, request *pb.SearchMessagesRequest) (*pb.SearchMessagesResponse, error) {
	messages, err := s.Messages.SearchMessages(ctx, request.DialogId, request.Query, request.OffsetId, request.Limit, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if messages == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no messages in dialog with id %s", request.DialogId))
	}
	response := new(pb.SearchMessagesResponse)
	response.Messages = messages
	return response, nil
}
