package user

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/user_v1"
)

type UserClient struct {
	clientUser user_v1.UserClient
}

func NewUserClient(client user_v1.UserClient) *UserClient {
	return &UserClient{
		clientUser: client,
	}
}

func (c *UserClient) Create(ctx context.Context, info *model.CreateUser) error {
	_, err := c.clientUser.Create(ctx, &user_v1.CreateRequest{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
		Role:            user_v1.Role(info.Role),
	})
	if err != nil {
		return err
	}

	return nil
}
