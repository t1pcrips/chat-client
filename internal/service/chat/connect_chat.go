package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/internal/service"
	"github.com/t1pcrips/chat-client/internal/service/utils"
)

func (s *ChatServiceImpl) ConnectChat(ctx context.Context, chatId int64, email string) (service.ServerStreamingClient, error) {
	tokens, user, err := s.getTokesnAndUser(ctx, email)
	if err != nil {
		return nil, err
	}

	newCtx := utils.NewOutgoingMetadataCtx(ctx, tokens.AccessToken)

	stream, err := s.clientChat.ConnectChat(newCtx, &model.ConnectChatRequest{
		ChatId:  chatId,
		Userame: user.Username,
	})
	if err != nil {
		return nil, err
	}

	return stream, nil
}
