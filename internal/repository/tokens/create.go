package tokens

import (
	"context"
	"encoding/json"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/internal/repository/tokens/model_tokens"
	"os"
)

func (repo *TokensRepositoryImpl) SaveToken(ctx context.Context, tokens *model.Tokens, email string) error {
	repoTokens := &model_tokens.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	file, err := os.OpenFile(repo.pathsConfig.FileTokensSave+email+".json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	err = json.NewEncoder(file).Encode(repoTokens)
	if err != nil {
		return err
	}

	return nil
}
