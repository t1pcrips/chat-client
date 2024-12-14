package auth

import (
	"github.com/t1pcrips/chat-client/pkg/auth_v1"
)

type AuthClient struct {
	clientAuth auth_v1.AuthClient
}

func NewAuthClient(client auth_v1.AuthClient) *AuthClient {
	return &AuthClient{
		clientAuth: client,
	}
}
