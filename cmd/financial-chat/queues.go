package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
	"github.com/streadway/amqp"
)

func configureQueues(conn *amqp.Connection, cmdQueueName string, responseQueueName string) (*queue.Consumer, *queue.Producer) {

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("error to open queue channel: %s", err.Error())
		return nil, nil
	}

	cmdsQueue, err := ch.QueueDeclare(
		"financial-cmds", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	responseQueue, err := ch.QueueDeclare(
		"financial-responses", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	producer := queue.Producer{
		Queue:   cmdsQueue,
		Channel: ch,
	}

	consumer := queue.Consumer{
		Queue:   responseQueue,
		Channel: ch,
	}

	return &consumer, &producer
}
