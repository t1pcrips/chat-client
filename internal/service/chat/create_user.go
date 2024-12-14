package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
)

func (s *ChatServiceImpl) CreateUser(ctx context.Context, info *model.CreateUser) error {
	err := s.clientUser.Create(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
