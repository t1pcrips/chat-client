package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/pkg/chat_v1"
)

func (c *ChatClient) Create(ctx context.Context, info []string) (int64, error) {
	chatId, err := c.clientChat.Create(ctx, &chat_v1.CreateRequest{
		Usernames: info,
	})
	if err != nil {
		return 0, err
	}

	return chatId.GetId(), nil
}
