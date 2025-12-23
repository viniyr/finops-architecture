package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"aevyn/finops-arch/internal/infrastructure/rabbitmq"
)

func main() {
	rabbitURL := "amqp://guest:guest@localhost:5672/"
	exchange := "x.disputes"

	queueName := "q.dispute.opened.card.v1"
	bindingKey := "dispute.opened.card.v1"

	conn, ch, err := rabbitmq.Dial(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	defer ch.Close()

	if err := rabbitmq.DeclareTopicExchange(ch, exchange); err != nil {
		log.Fatal(err)
	}

	if err := ch.Qos(10, 0, false); err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := ch.QueueBind(
		q.Name,
		bindingKey,
		exchange,
		false,
		nil,
	); err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"processor-card-v1",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("consumer started queue=%s binding=%s exchange=%s", q.Name, bindingKey, exchange)

	for {
		select {
		case <-ctx.Done():
			log.Println("shutdown consumer")
			return

		case d, ok := <-msgs:
			if !ok {
				log.Println("delivery ch closed")
				return
			}

			if err := handleDisputeOpenedCard(d.Body); err != nil {
				_ = d.Nack(false, true)
				continue
			}

			_ = d.Ack(false)
		}
	}
}

func handleDisputeOpenedCard(body []byte) error {
	log.Printf("received %s", string(body))

	time.Sleep(50 * time.Millisecond)
	return nil
}
