package concurrency

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string) // 2 Channels vom Type string
	c2 := make(chan string)

	go func() { // Funktion, die in den 1. Channel schreibt
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() { // Funktion, die in den 2. Channel schreibt
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ { // Schleife, die den select ausführt und den case ausführt, der gerade etwas zu Lesen bekommt
		select {
		case msg1 := <-c1: // hier wird versucht vom 1. Channel zu lesen, wenn es nichts gibt, geht es in den nächsten case
			fmt.Println("Received", msg1)
		case msg2 := <-c2: // hier wird versucht vom 2. Channel zu lesen, wenn es nichts gibt, geht es in den nächsten case
			fmt.Println("Received", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: no message received after 3s")
		}
	}

}
