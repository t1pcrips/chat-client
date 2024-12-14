package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/internal/service/utils"
	"time"
)

func (s *ChatServiceImpl) SendMessage(ctx context.Context, chatId int64, text string, timeSend time.Time) error {
	tokens, claims, err := s.getTokesnAndUser(ctx, "")
	if err != nil {
		return err
	}

	newCtx := utils.NewOutgoingMetadataCtx(ctx, tokens.AccessToken)

	err = s.clientChat.SendMessage(newCtx, &model.Message{
		ToChatId:  chatId,
		From:      claims.Username,
		Text:      text,
		Timestamp: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}
