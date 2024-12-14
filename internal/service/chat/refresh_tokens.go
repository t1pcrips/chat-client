package chat

import "context"

func (s *ChatServiceImpl) RefreshTokens(ctx context.Context, token string) error {
	tokens, err := s.clientAuth.RefreshTokens(ctx, token)
	if err != nil {
		return err
	}

	err = s.tokensRepository.SaveToken(ctx, tokens, "")
	if err != nil {
		return err
	}

	return nil
}
