package gobasic

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	msg1 := "hi1"
	//msg2 := "hi2"

	go func() {
		//time.Sleep(2 * time.Second)
		messages <- msg1
	}()

	time.Sleep(2 * time.Second)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message3", msg)
	case sig := <-signals:
		fmt.Println("received signal3", sig)
	default:
		fmt.Println("no activity3")
	}
}
