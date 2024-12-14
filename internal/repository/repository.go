package repository

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
)

type TokensRepository interface {
	SaveToken(ctx context.Context, tokens *model.Tokens) error
	GetToken(ctx context.Context) (*model.Tokens, error)
}
