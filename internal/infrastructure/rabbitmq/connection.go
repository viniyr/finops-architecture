package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Dial(url string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}
