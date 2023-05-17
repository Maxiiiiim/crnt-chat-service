package auth

import (
	"context"
	"google.golang.org/grpc"
)

func (a *AuthMiddleware) GetInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		payload, err := a.authorizer.AuthorizeUser(ctx)
		if err != nil {
			return nil, err
		}
		newCtx := context.WithValue(ctx, "auth", payload)

		return handler(newCtx, req)
	}
}
