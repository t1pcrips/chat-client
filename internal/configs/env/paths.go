package env

import (
	"errors"
	"github.com/t1pcrips/chat-client/internal/configs"
	"os"
)

const (
	fileTokensSave = "FILE_TOKENS_SAVE"
)

type PathsConfigSearcher struct{}

func NewPathsConfigSearcher() *PathsConfigSearcher {
	return &PathsConfigSearcher{}
}

func (cfg *PathsConfigSearcher) Get() (*configs.PathsConfig, error) {
	fileTokensSave := os.Getenv(fileTokensSave)
	if len(fileTokensSave) == 0 {
		return nil, errors.New("file tokens save not found")
	}

	return &configs.PathsConfig{
		FileTokensSave: fileTokensSave,
	}, nil
}
