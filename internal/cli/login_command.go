package cli

import (
	"github.com/spf13/cobra"
	"log"
)

func (c *Chat) Login() *cobra.Command {
	return &cobra.Command{
		Use:   login,
		Short: loginDesc,
		Run: func(cmd *cobra.Command, args []string) {
			email, err := cmd.Flags().GetString(email)
			if err != nil {
				log.Println(err.Error())
				return
			}

			password, err := cmd.Flags().GetString(password)
			if err != nil {
				log.Println(err.Error())
				return
			}

			err = c.chatService.LoginUser(cmd.Context(), email, password)
			if err != nil {
				log.Println(err.Error())
				return
			}

			c.writer.Info("login successful")
		},
	}
}
