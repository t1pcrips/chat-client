package cli

import "time"

type ConsoleWriter interface {
	Info(msg string)
	Error(msg string)
	Message(msgDateTime time.Time, userName string, msg string)
	ScanMessage() (string, error)
	CleanUpLine()
}
