package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"net"
)

type PathsConfig struct {
	FileTokensSave string
}

type AuthClientConfig struct {
	Host string
	Port string
}

type UserClientConfig struct {
	Host string
	Port string
}

type ChatClientConfig struct {
	Host string
	Port string
}

func (cfg *AuthClientConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *UserClientConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *ChatClientConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return errors.New("faild to load (...).env file")
	}

	return nil
}
