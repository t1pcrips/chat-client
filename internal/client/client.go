package client

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
)

type AuthClient interface {
	Login(ctx context.Context, info *model.LoginAuthRequest) (*model.Tokens, error)
	RefreshTokens(ctx context.Context, token string) (*model.Tokens, error)
}

type ChatClient interface {
	Create(ctx context.Context, info []string) (int64, error)
	SendMessage(ctx context.Context, info *model.Message) error
	ConnectChat(ctx context.Context, info *model.ConnectChatRequest) (ServerStreamingClient, error)
}

type UserClient interface {
	Create(ctx context.Context, info *model.CreateUser) error
}

type ServerStreamingClient interface {
	Recv() (*model.Message, error)
	Context() context.Context
}
