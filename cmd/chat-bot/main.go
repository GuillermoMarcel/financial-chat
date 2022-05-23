package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
	"github.com/streadway/amqp"
)

func main() {
	logger := log.Default()

	//Queue
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Printf("error to connecto to queue %s", err.Error())
		return
	}
	defer conn.Close()
	logger.Println("queue connected")

	ch, err := conn.Channel()
	if err != nil {
		logger.Printf("2 error to connecto to queue %s", err.Error())
		return
	}
	defer ch.Close()

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
	consumer := queue.Consumer{
		Queue:   cmdsQueue,
		Channel: ch,
	}

	producer := queue.Producer{
		Queue:   responseQueue,
		Channel: ch,
	}

	defer consumer.Stop()

	producer.SendMessage("que se yo")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever

}
