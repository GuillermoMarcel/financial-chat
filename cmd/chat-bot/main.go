package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/chat-bot/controller"
	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
	"github.com/streadway/amqp"
)

func main() {
	logger := log.Default()

	config := readConfigs(logger)
	if config == nil {
		return
	}

	//Queue
	conn, err := amqp.Dial(config.BrokerAddress)
	if err != nil {
		logger.Printf("error to connecto to queue %s", err.Error())
		return
	}
	defer conn.Close()
	logger.Println("queue connected")

	ch, err := conn.Channel()
	if err != nil {
		logger.Printf("error opening channel %s", err.Error())
		return
	}
	defer ch.Close()

	cmdsQueue, err := ch.QueueDeclare(
		config.CmdQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	responseQueue, err := ch.QueueDeclare(
		config.ResponsesQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	consumer := &queue.AmqpConsumer{
		Queue:   cmdsQueue,
		Channel: ch,
	}

	producer := queue.Producer{
		Queue:   responseQueue,
		Channel: ch,
	}
	defer consumer.Stop()

	controller := controller.RequestController{
		CmdConsumer:    consumer,
		ReturnProducer: producer,
	}

	controller.ServeApp()

}
