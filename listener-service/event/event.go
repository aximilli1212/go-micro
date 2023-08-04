package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      //topic
		true,         //durable?
		false,        // auto-deleted?
		false,        //internal?
		false,        // nowait?
		nil,          // arguments?
	)
}
