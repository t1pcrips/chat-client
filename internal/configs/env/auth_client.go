package env

import (
	"errors"
	"github.com/t1pcrips/chat-client/internal/configs"
	"os"
	"strconv"
)

const (
	authPort = "AUTH_PORT"
	authHost = "AUTH_HOST"
)

type AuthConfigSearcher struct{}

func NewAuthConfigSearcher() *AuthConfigSearcher {
	return &AuthConfigSearcher{}
}

func (cfg *AuthConfigSearcher) Get() (*configs.AuthClientConfig, error) {
	host := os.Getenv(authHost)
	if len(host) == 0 {
		return nil, errors.New("auth host not found")
	}

	portString := os.Getenv(authPort)
	if len(host) == 0 {
		return nil, errors.New("auth port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("user integer port")
	}

	return &configs.AuthClientConfig{
		Host: host,
		Port: portString,
	}, nil
}
