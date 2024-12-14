package env

import (
	"errors"
	"github.com/t1pcrips/chat-client/internal/configs"
	"os"
	"strconv"
)

const (
	chatPort = "CHAT_PORT"
	chatHost = "CHAT_HOST"
)

type ChatConfigSearcher struct{}

func NewChatConfigSearcher() *ChatConfigSearcher {
	return &ChatConfigSearcher{}
}

func (cfg *ChatConfigSearcher) Get() (*configs.ChatClientConfig, error) {
	host := os.Getenv(chatHost)
	if len(host) == 0 {
		return nil, errors.New("chat host not found")
	}

	portString := os.Getenv(chatPort)
	if len(host) == 0 {
		return nil, errors.New("chat port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("user integer port")
	}

	return &configs.ChatClientConfig{
		Host: host,
		Port: portString,
	}, nil
}
