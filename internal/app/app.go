package app

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/cli"
	"github.com/t1pcrips/chat-client/internal/configs"
	"github.com/t1pcrips/platform-pkg/pkg/closer"
)

type App struct {
	serviceProvider *serviceProvider
	cli             *cli.Chat
	configPath      string
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.cli.Execute(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initCli,
	}

	for _, f := range deps {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := configs.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initCli(ctx context.Context) error {
	a.cli = cli.NewChat(
		a.serviceProvider.ChatService(),
		a.serviceProvider.Writer(),
	)

	return nil
}
