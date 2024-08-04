package main

import (
	"fmt"
	"time"
)

type Message struct {
	id      int64
	message string
}

var messageChannel = make(chan Message)

func (m *Message) WriteMessage(msg string) {
	m.id = time.Now().UnixMilli()
	m.message = msg
}

func readMessageChannel() {
	for {
		msg := <-messageChannel
		fmt.Println(msg)
	}
}

func main() {
	var msg Message

	go readMessageChannel()

	for {
		var message string

		fmt.Println("Enter a message: ")
		_, err := fmt.Scanf("%s", &message)

		if err != nil {
			fmt.Println("Error reading string: ", err)
			continue
		}

		msg.WriteMessage(message)

		messageChannel <- msg

	}
}
