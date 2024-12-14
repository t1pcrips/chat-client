package tokens

import (
	"context"
	"encoding/json"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/internal/repository/tokens/model_tokens"
	"os"
)

func (repo *TokensRepositoryImpl) GetToken(ctx context.Context) (*model.Tokens, error) {
	repoTokens := &model_tokens.Tokens{}

	file, err := os.OpenFile(repo.pathsConfig.FileTokensSave, os.O_RDONLY, 1111)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	err = json.NewDecoder(file).Decode(repoTokens)
	if err != nil {
		return nil, err
	}

	tokens := &model.Tokens{
		RefreshToken: repoTokens.RefreshToken,
		AccessToken:  repoTokens.AccessToken,
	}

	return tokens, nil
}
