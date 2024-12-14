package app

import (
	"github.com/t1pcrips/chat-client/internal/cli"
	"github.com/t1pcrips/chat-client/internal/cli/console"
	"github.com/t1pcrips/chat-client/internal/client"
	"github.com/t1pcrips/chat-client/internal/client/auth"
	"github.com/t1pcrips/chat-client/internal/client/chat"
	"github.com/t1pcrips/chat-client/internal/client/user"
	"github.com/t1pcrips/chat-client/internal/configs"
	"github.com/t1pcrips/chat-client/internal/configs/env"
	"github.com/t1pcrips/chat-client/internal/repository"
	"github.com/t1pcrips/chat-client/internal/repository/tokens"
	"github.com/t1pcrips/chat-client/internal/service"
	chatService "github.com/t1pcrips/chat-client/internal/service/chat"
	"github.com/t1pcrips/chat-client/pkg/auth_v1"
	"github.com/t1pcrips/chat-client/pkg/chat_v1"
	"github.com/t1pcrips/chat-client/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type serviceProvider struct {
	pathsConfig      *configs.PathsConfig
	authClientConfig *configs.AuthClientConfig
	userClientConfig *configs.UserClientConfig
	chatClientConfig *configs.ChatClientConfig

	authV1Client auth_v1.AuthClient
	chatV1Client chat_v1.ChatClient
	userV1Client user_v1.UserClient
	//streamingClient chat_v1.Chat_ConnectClient

	authClient   client.AuthClient
	chatClient   client.ChatClient
	userClient   client.UserClient
	streamClient client.ServerStreamingClient

	chatService service.ChatService

	writer cli.ConsoleWriter

	tokensRepository repository.TokensRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Writer() cli.ConsoleWriter {
	if s.writer == nil {
		s.writer = console.NewConsoleWriter()
	}

	return s.writer
}

func (s *serviceProvider) PathConfig() *configs.PathsConfig {
	if s.pathsConfig == nil {
		cfgSearcher := env.NewPathsConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get path config: %s", err.Error())
		}

		s.pathsConfig = cfg
	}

	return s.pathsConfig
}

func (s *serviceProvider) AuthClientConfig() *configs.AuthClientConfig {
	if s.authClientConfig == nil {
		cfgSearcher := env.NewAuthConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authClientConfig = cfg
	}

	return s.authClientConfig
}

func (s *serviceProvider) UserClientConfig() *configs.UserClientConfig {
	if s.userClientConfig == nil {
		cfgSearcher := env.NewUserConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get user config: %s", err.Error())
		}

		s.userClientConfig = cfg
	}

	return s.userClientConfig
}

func (s *serviceProvider) ChatClientConfig() *configs.ChatClientConfig {
	if s.chatClientConfig == nil {
		cfgSearcher := env.NewChatConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get chat config: %s", err.Error())
		}

		s.chatClientConfig = cfg
	}

	return s.chatClientConfig
}

func (s *serviceProvider) ConsoleWriter() cli.ConsoleWriter {
	if s.writer == nil {
		s.writer = console.NewConsoleWriter()
	}

	return s.writer
}

func (s *serviceProvider) AuthV1Client() auth_v1.AuthClient {
	conn, err := grpc.NewClient(
		s.AuthClientConfig().Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err.Error())
	}

	return auth_v1.NewAuthClient(conn)
}

func (s *serviceProvider) ChatV1Client() chat_v1.ChatClient {
	conn, err := grpc.NewClient(
		s.ChatClientConfig().Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err.Error())
	}

	return chat_v1.NewChatClient(conn)
}

func (s *serviceProvider) UserV1Client() user_v1.UserClient {
	conn, err := grpc.NewClient(
		s.UserClientConfig().Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to user service: %s", err.Error())
	}

	return user_v1.NewUserClient(conn)
}

func (s *serviceProvider) AuthClient() client.AuthClient {
	if s.authClient == nil {
		s.authClient = auth.NewAuthClient(s.AuthV1Client())
	}

	return s.authClient
}

func (s *serviceProvider) ChatClient() client.ChatClient {
	if s.chatClient == nil {
		s.chatClient = chat.NewChatClient(s.ChatV1Client())
	}

	return s.chatClient
}

func (s *serviceProvider) UserClient() client.UserClient {
	if s.userClient == nil {
		s.userClient = user.NewUserClient(s.UserV1Client())
	}

	return s.userClient
}

func (s *serviceProvider) ChatService() service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatServiceImpl(
			s.UserClient(),
			s.ChatClient(),
			s.AuthClient(),
			s.TokensRepository(),
		)
	}

	return s.chatService
}

//func (s *serviceProvider) ServerStreamingClient() client.ServerStreamingClient {
//	if s.streamClient == nil {
//		s.streamClient = chat.NewServerStreamingClient()
//	}
//
//	return s.streamClient
//}

func (s *serviceProvider) TokensRepository() repository.TokensRepository {
	if s.tokensRepository == nil {
		s.tokensRepository = tokens.NewTokensRepositoryImpl(s.PathConfig())
	}

	return s.tokensRepository
}
