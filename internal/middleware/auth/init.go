package auth

import (
	sdk_auth "github.com/Constantine27K/crnt-sdk/pkg/authorization"
	sdk_token "github.com/Constantine27K/crnt-sdk/pkg/token"
	"log"
	"os"
)

type AuthMiddleware struct {
	authorizer sdk_auth.Authorizer
}

func NewAuthMiddleware() *AuthMiddleware {
	key := os.Getenv("SYMMETRIC_KEY")
	maker, err := sdk_token.NewMaker(key)
	if err != nil {
		log.Fatal(err)
	}

	auth := sdk_auth.NewAuthorizer(maker)
	return &AuthMiddleware{auth}
}
