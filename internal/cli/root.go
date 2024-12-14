package cli

import (
	"github.com/spf13/cobra"
	"github.com/t1pcrips/chat-client/internal/service"
	"log"
)

const (
	appName         = "chat-project"
	appDesc         = "cli util for chat"
	create          = "create"
	createDesc      = "create user or chat"
	user            = "user"
	userDesc        = "create user"
	login           = "login"
	loginDesc       = "login user"
	chat            = "chat"
	chatDesc        = "create chat for users"
	connect         = "connect-chat"
	connectDesc     = "connecting to created chat"
	username        = "username"
	email           = "email"
	password        = "password"
	passwordConfirm = "password-confirm"
	chatId          = "chat-id"
	role            = "role"
)

type Chat struct {
	chatService service.ChatService

	rootCommand       *cobra.Command
	createCommand     *cobra.Command
	chatCommand       *cobra.Command
	loginCommand      *cobra.Command
	connectCommand    *cobra.Command
	createUserCommand *cobra.Command
	createChatCommand *cobra.Command

	writer ConsoleWriter
}

// chat-project create user -u hiak -e tima@gaf.com -p 123123123 -pc 123123123 -r admin
// chat-project create chat -u qwe,rty
// chat-project connect -u qwe -c 1

func NewChat(chatService service.ChatService, writer ConsoleWriter) *Chat {
	return &Chat{
		chatService: chatService,
		writer:      writer,
	}
}

func (c *Chat) createRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   appName,
		Short: appDesc,
	}
}

func (c *Chat) createCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   create,
		Short: createDesc,
	}
}

func (c *Chat) commands() {
	c.rootCommand.AddCommand(c.createCommand)
	c.rootCommand.AddCommand(c.chatCommand)
	c.rootCommand.AddCommand(c.connectCommand)

	c.createCommand.AddCommand(c.createUserCommand)
	c.createCommand.AddCommand(c.createChatCommand)

	c.createUserCommand.Flags().StringP(username, "u", "", "name of user")
	c.createUserCommand.Flags().StringP(email, "e", "", "email of user")
	c.createUserCommand.Flags().StringP(password, "p", "", "password of user")
	c.createUserCommand.Flags().StringP(passwordConfirm, "pc", "", "passwords confirm for user")
	c.createUserCommand.MarkFlagsRequiredTogether(username, email, password, passwordConfirm)

	c.loginCommand.Flags().StringP(email, "e", "", "email of user")
	c.loginCommand.Flags().StringP(password, "p", "", "password of user")
	c.loginCommand.MarkFlagsRequiredTogether(username, email)

	c.connectCommand.Flags().Int64P(chatId, "c", 0, "chat id for connecting")
	err := c.connectCommand.MarkFlagRequired(chatId)
	if err != nil {
		log.Println(err.Error())
	}
}
