package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Consumer struct {
	Queue      amqp.Queue
	Channel    *amqp.Channel
	endChan    chan bool
	ReturnChan chan []byte
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
				fmt.Printf("message recived, queue:%s\n", c.Queue.Name)
				c.ReturnChan <- d.Body
			case <-c.endChan:
				close(c.endChan)
				c.endChan = nil
				return
			}
		}
	}()
}

func (c Consumer) Stop() {
	if c.endChan != nil {
		c.endChan <- true
	}
	c.Channel.Close()
}
