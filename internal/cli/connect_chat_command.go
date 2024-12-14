package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/t1pcrips/chat-client/internal/service"
	"io"
	"log"
	"os"
	"time"
)

func (c *Chat) createConnectCommand() *cobra.Command {
	return &cobra.Command{
		Use:   connect,
		Short: connectDesc,
		Run: func(cmd *cobra.Command, args []string) {
			chatId, err := cmd.Flags().GetInt64(chatId)
			if err != nil {
				log.Println(err.Error())
				return
			}

			email, err := cmd.Flags().GetString(email)
			if err != nil {
				log.Println(err.Error())
				return
			}

			stream, err := c.chatService.ConnectChat(cmd.Context(), chatId, email)
			if err != nil {
				c.writer.Error(err.Error())
				return
			}

			c.writer.Info("successful to connect chat")

			go func() {
				c.incomingMessages(stream)
			}()

			c.outgoingMessages(cmd.Context(), chatId)
		},
	}
}

func (c *Chat) incomingMessages(stream service.ServerStreamingClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return
		}

		if err != nil {
			c.writer.Error(err.Error())
			return
		}
		c.writer.Message(msg.Timestamp, msg.From, msg.Text)
	}
}

func (c *Chat) outgoingMessages(ctx context.Context, chatId int64) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		msg, err := c.writer.ScanMessage()
		if err != nil {
			c.writer.Info("exit chat")
			break
		}

		c.writer.CleanUpLine()

		err = c.chatService.SendMessage(ctx, chatId, msg, time.Now())
		if err != nil {
			fmt.Println(err)
			c.writer.Error("failed to  connect chat")
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Println("failed to scan messages: ", err.Error())
	}
}
