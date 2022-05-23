package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
}

func (p Producer) SendJson(message interface{}) {

	content, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        content,
		})
	if err != nil {
		fmt.Printf("couldn't send message: %s\n", err.Error())
		return
	}

	fmt.Println("mensaje enviado")
}
