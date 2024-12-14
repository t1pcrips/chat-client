package auth

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/auth_v1"
)

func (c *AuthClient) RefreshTokens(ctx context.Context, token string) (*model.Tokens, error) {
	resp, err := c.clientAuth.RefreshTokens(ctx, &auth_v1.RefreshTokensRequest{
		RefreshToken: token,
	})
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		RefreshToken: resp.GetRefreshToken(),
		AccessToken:  resp.GetAccessToken(),
	}, nil
}
