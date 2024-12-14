package configs

import (
	"errors"
	"github.com/joho/godotenv"
)

type PathsConfig struct {
	FileTokensSave string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return errors.New("faild to load (...).env file")
	}

	return nil
}
