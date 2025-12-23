package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Ch       *amqp.Channel
	Exchange string
}

func (p Publisher) Publish(ctx context.Context, routingKey string, body []byte) error {
	return p.Ch.PublishWithContext(
		ctx,
		p.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
