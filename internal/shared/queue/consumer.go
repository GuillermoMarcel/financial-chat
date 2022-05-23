package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Consumer interface {
	Start()
	Stop()
	RegisterReturnChan(chan []byte)
}

type AmqpConsumer struct {
	Queue      amqp.Queue
	Channel    *amqp.Channel
	endChan    chan bool
	ReturnChan chan []byte
}

func (c *AmqpConsumer) Start() {
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

	closeChan := make(chan *amqp.Error)
	c.Channel.NotifyClose(closeChan)

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

			case closeMsg := <-closeChan:
				//Shoud alert to close or retry connection.
				fmt.Printf("channel closed: %s\n", closeMsg.Reason)
				return
			}
		}
	}()
}

func (c *AmqpConsumer) Stop() {
	if c.endChan != nil {
		c.endChan <- true
	}
	c.Channel.Close()
	c.ReturnChan = nil
}

func (c *AmqpConsumer) RegisterReturnChan(ch chan []byte) {
	c.ReturnChan = ch
}
