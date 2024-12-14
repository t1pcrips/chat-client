package console

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"sync"
	"time"
)

const (
	format = "2006-01-02 15:04:05"
)

type Writer struct {
	mx      sync.Mutex
	scanner *bufio.Scanner
}

func NewConsoleWriter() *Writer {
	return &Writer{
		mx:      sync.Mutex{},
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *Writer) Info(text string) {
	printMsg := color.YellowString(fmt.Sprintf("[%v]", time.Now().Format(format)))
	printMsg += color.GreenString(" INFO -")
	printMsg += text
	fmt.Println(printMsg)
}

func (c *Writer) Error(text string) {
	printMsg := color.YellowString(fmt.Sprintf("[%v]", time.Now().Format(format)))
	printMsg += color.RedString("ERROR - ")
	printMsg += text
	fmt.Println(printMsg)
}

func (c *Writer) Message(msgDateTime time.Time, userName string, msg string) {
	printMsg := color.YellowString(fmt.Sprintf("[%v]", time.Now().Format(format)))
	printMsg += color.CyanString(fmt.Sprintf("FROM: %s", userName))
	printMsg += msg
	fmt.Println(printMsg)
}

func (c *Writer) ScanMessage() (string, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.scanner.Scan()

	msg := c.scanner.Text()
	if msg == "exit" {
		return "", errors.New("exit chat")
	}

	return msg, nil
}

func (c *Writer) CleanUpLine() {
	c.mx.Lock()
	defer c.mx.Unlock()
	fmt.Printf("\033[1A\033[K")
}
