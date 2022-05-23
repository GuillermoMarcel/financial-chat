package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
}

func (p Producer) SendMessage(message string) {

	err := p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		fmt.Printf("3 mensaje no enviado: %s\n", err.Error())
		return
	}

	fmt.Println("mensaje enviado")
}
