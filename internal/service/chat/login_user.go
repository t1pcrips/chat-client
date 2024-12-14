package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
)

func (s *ChatServiceImpl) LoginUser(ctx context.Context, email string, password string) error {
	tokens, err := s.clientAuth.Login(ctx, &model.LoginAuthRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return s.tokensRepository.SaveToken(ctx, tokens)
}
