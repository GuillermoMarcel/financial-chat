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
		fmt.Printf("queue unable to start consuming: %s\n", err.Error())
		return
	}

	c.endChan = make(chan bool)

	go func() {
		fmt.Println("start reading queue messages")
		for {
			select {
			case d := <-msgs:
				fmt.Printf("consumer: message recived, queue:%s\n", c.Queue.Name)
				if c.ReturnChan == nil {
					fmt.Println("consumer return channel not set, mesage lost")
					continue
				}
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
