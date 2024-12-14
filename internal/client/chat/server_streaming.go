package chat

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/model"
	"github.com/t1pcrips/chat-client/pkg/chat_v1"
)

type ServerStreamingClient struct {
	streamingClient chat_v1.Chat_ConnectClient
}

func NewServerStreamingClient(streamingClient chat_v1.Chat_ConnectClient) *ServerStreamingClient {
	return &ServerStreamingClient{
		streamingClient: streamingClient,
	}
}

func (s *ServerStreamingClient) Recv() (*model.Message, error) {
	message, err := s.streamingClient.Recv()
	if err != nil {
		return nil, err
	}

	return &model.Message{
		From:      message.GetFrom(),
		Text:      message.GetText(),
		Timestamp: message.GetTimestamp().AsTime(),
		ToChatId:  message.GetToChatId(),
	}, nil
}

func (s *ServerStreamingClient) Context() context.Context {
	return s.streamingClient.Context()
}
