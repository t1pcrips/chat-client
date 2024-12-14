package tokens

import (
	"github.com/t1pcrips/chat-client/internal/configs"
	"github.com/t1pcrips/chat-client/internal/repository"
)

type TokensRepositoryImpl struct {
	pathsConfig *configs.PathsConfig
}

func NewTokensRepositoryImpl(pathsConfig *configs.PathsConfig) repository.TokensRepository {
	return &TokensRepositoryImpl{
		pathsConfig: pathsConfig,
	}
}
