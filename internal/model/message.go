package model

import (
	"time"
)

type Message struct {
	From      string
	Text      string
	Timestamp time.Time
	ToChatId  int64
}
