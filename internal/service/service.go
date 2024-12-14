package service

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"time"
)

type ChatService interface {
	CreateUser(ctx context.Context, info *model.CreateUser) error
	LoginUser(ctx context.Context, email string, password string) error
	RefreshTokens(ctx context.Context, token string) error
	CreateChat(ctx context.Context, email string) (int64, error)
	ConnectChat(ctx context.Context, chatId int64, email string) (ServerStreamingClient, error)
	SendMessage(ctx context.Context, chatId int64, text string, timeSend time.Time) error
}

type ServerStreamingClient interface {
	Recv() (*model.Message, error)
	Context() context.Context
}
