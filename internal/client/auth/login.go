package auth

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/auth_v1"
)

func (c *AuthClient) Login(ctx context.Context, info *model.LoginAuthRequest) (*model.Tokens, error) {
	resp, err := c.clientAuth.Login(ctx, &auth_v1.LoginRequest{
		Email:    info.Email,
		Password: info.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		RefreshToken: resp.GetRefreshToken(),
		AccessToken:  resp.GetAccessToken(),
	}, nil
}
