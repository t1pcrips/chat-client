package repository

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
)

type TokensRepository interface {
	SaveToken(ctx context.Context, tokens *model.Tokens, email string) error
	GetToken(ctx context.Context, email string) (*model.Tokens, error)
}
