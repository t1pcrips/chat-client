package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/client"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/chat_v1"
)

func (c *ChatClient) ConnectChat(ctx context.Context, info *model.ConnectChatRequest) (client.ServerStreamingClient, error) {
	stream, err := c.clientChat.Connect(ctx, &chat_v1.ConnectChatRequest{
		ChatId:   info.ChatId,
		Username: info.Userame,
	})
	if err != nil {
		return nil, err
	}

	return NewServerStreamingClient(stream), nil
}
