package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *ChatClient) SendMessage(ctx context.Context, info *model.Message) error {
	_, err := c.clientChat.SendMessage(ctx, &chat_v1.SendMessageRequest{
		From:      info.From,
		Text:      info.Text,
		ToChatId:  info.ToChatId,
		Timestamp: timestamppb.New(info.Timestamp),
	})
	if err != nil {
		return err
	}

	return nil
}
