package chat

import (
	"github.com/t1pcrips/chat-client/internal/client"
	"github.com/t1pcrips/chat-client/internal/repository"
	"github.com/t1pcrips/chat-client/internal/service"
)

type ChatServiceImpl struct {
	clientUser client.UserClient
	clientChat client.ChatClient
	clientAuth client.AuthClient

	tokensRepository repository.TokensRepository
}

func NewChatServiceImpl(
	clientUser client.UserClient,
	clientChat client.ChatClient,
	clientAuth client.AuthClient,
	tokensRepository repository.TokensRepository,
) service.ChatService {
	return &ChatServiceImpl{
		clientUser:       clientUser,
		clientAuth:       clientAuth,
		clientChat:       clientChat,
		tokensRepository: tokensRepository,
	}
}
