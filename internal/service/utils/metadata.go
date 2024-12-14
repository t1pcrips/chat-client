package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const (
	authorization = "Authorization"
	bearer        = "Bearer "
)

func NewOutgoingMetadataCtx(ctx context.Context, token string) context.Context {
	outgoingMd := metadata.New(map[string]string{authorization: bearer + token})
	return metadata.NewOutgoingContext(ctx, outgoingMd)
}
