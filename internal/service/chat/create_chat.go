package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/service/utils"
)

func (s *ChatServiceImpl) CreateChat(ctx context.Context, email string) (int64, error) {
	tokens, user, err := s.getTokesnAndUser(ctx, email)
	if err != nil {
		return 0, err
	}

	newCtx := utils.NewOutgoingMetadataCtx(ctx, tokens.AccessToken)

	chatId, err := s.clientChat.Create(newCtx, []string{user.Username})
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
