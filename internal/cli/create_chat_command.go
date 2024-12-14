package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func (c *Chat) createCreateChatCommand() *cobra.Command {
	return &cobra.Command{
		Use:   chat,
		Short: chatDesc,
		Run: func(cmd *cobra.Command, args []string) {
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				log.Println(err.Error())
				return
			}

			chatId, err := c.chatService.CreateChat(cmd.Context(), email)
			if err != nil {
				log.Println(err.Error())
				return
			}
			c.writer.Info(fmt.Sprintf("chat created with id: %d", chatId))
		},
	}
}
