package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"frodo@underhill.me"},
		From:    "gandalf@whilecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	FailedMessage := FailedMessage{
		ErrorMessage:    "Message intercepted by black rider",
		OriginalMessage: Message{},
	}
	msgCh <- msg
	errCh <- FailedMessage

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
