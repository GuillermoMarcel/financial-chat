package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	endChan chan bool
}

func (c Consumer) Start() {
	msgs, err := c.Channel.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		fmt.Printf("3 mensaje no enviado: %s\n", err.Error())
		return
	}

	c.endChan = make(chan bool)

	go func() {
		for {
			select {
			case d := <-msgs:
				log.Printf("Received a message: %s", d.Body)
			case <-c.endChan:
				return
			}
		}
	}()
}

func (c Consumer) Stop() {
	c.endChan <- true
	c.Channel.Close()
}
