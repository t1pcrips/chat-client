package chat

import "github.com/t1pcrips/chat-client/pkg/chat_v1"

type ChatClient struct {
	clientChat chat_v1.ChatClient
}

func NewChatClient(client chat_v1.ChatClient) *ChatClient {
	return &ChatClient{
		clientChat: client,
	}
}
