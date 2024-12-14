package cli

import (
	"github.com/spf13/cobra"
	"github.com/t1pcrips/chat-client/internal/model"
	"log"
)

func (c *Chat) createCreateUserCommand() *cobra.Command {
	return &cobra.Command{
		Use:   user,
		Short: userDesc,
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString(username)
			if err != nil {
				log.Println(err.Error())
				return
			}

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

			passwordConf, err := cmd.Flags().GetString(passwordConfirm)
			if err != nil {
				log.Println(err.Error())
				return
			}

			role, err := cmd.Flags().GetInt64(role)
			if err != nil {
				log.Println(err.Error())
				return
			}

			err = c.chatService.CreateUser(cmd.Context(), &model.CreateUser{
				Name:            username,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConf,
				Role:            role,
			})

			if err != nil {
				log.Println(err.Error())
				return
			}

			c.writer.Info("user successful created")
		},
	}
}
