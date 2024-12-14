package chat

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/t1pcrips/chat-client/internal/model"
)

func (s *ChatServiceImpl) getTokesnAndUser(ctx context.Context, email string) (*model.Tokens, *model.UserClaims, error) {
	tokens, err := s.tokensRepository.GetToken(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokens.AccessToken, &model.UserClaims{})
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, nil, errors.New("failed to get claims")
	}

	return tokens, claims, nil
}
