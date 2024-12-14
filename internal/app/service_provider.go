package app

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/configs"
	"github.com/t1pcrips/chat-client/internal/configs/env"
	"github.com/t1pcrips/chat-client/internal/repository"
	"github.com/t1pcrips/chat-client/internal/repository/tokens"
	"log"
)

type serviceProvider struct {
	pathsConfig *configs.PathsConfig

	tokensRepository repository.TokensRepository
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PathConfig() *configs.PathsConfig {
	if s.pathsConfig == nil {
		cfgSearcher := env.NewPathsConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get path config: %s", err.Error())
		}

		s.pathsConfig = cfg
	}

	return s.pathsConfig
}

func (s *serviceProvider) TokensRepository(ctx context.Context) repository.TokensRepository {
	if s.tokensRepository == nil {
		s.tokensRepository = tokens.NewTokensRepositoryImpl(s.PathConfig())
	}

	return s.tokensRepository
}
