package env

import (
	"errors"
	"github.com/t1pcrips/chat-client/internal/configs"
	"os"
	"strconv"
)

const (
	userPort = "USER_PORT"
	userHost = "USER_HOST"
)

type UserConfigSearcher struct{}

func NewUserConfigSearcher() *UserConfigSearcher {
	return &UserConfigSearcher{}
}

func (cfg *UserConfigSearcher) Get() (*configs.UserClientConfig, error) {
	host := os.Getenv(userHost)
	if len(host) == 0 {
		return nil, errors.New("user host not found")
	}

	portString := os.Getenv(userPort)
	if len(host) == 0 {
		return nil, errors.New("user port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("user integer port")
	}

	return &configs.UserClientConfig{
		Host: host,
		Port: portString,
	}, nil
}
