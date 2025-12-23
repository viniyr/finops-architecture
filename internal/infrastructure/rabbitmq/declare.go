package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func DeclareTopicExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
