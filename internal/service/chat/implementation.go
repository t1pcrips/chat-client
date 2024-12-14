package chat

import (
	"github.com/t1pcrips/chat-client/internal/client/auth"
	"github.com/t1pcrips/chat-client/internal/client/chat"
	"github.com/t1pcrips/chat-client/internal/client/user"
	"github.com/t1pcrips/chat-client/internal/repository"
	"github.com/t1pcrips/chat-client/internal/service"
)

type ChatServiceImpl struct {
	clientUser user.UserClient
	clientChat chat.ChatClient
	clientAuth auth.AuthClient

	tokensRepository repository.TokensRepository
}

func NewChatServiceImpl(
	clientUser user.UserClient,
	clientChat chat.ChatClient,
	clientAuth auth.AuthClient,
	tokensRepository repository.TokensRepository,
) service.ChatService {
	return &ChatServiceImpl{
		clientUser:       clientUser,
		clientAuth:       clientAuth,
		clientChat:       clientChat,
		tokensRepository: tokensRepository,
	}
}
