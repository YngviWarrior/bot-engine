package rabbitmq

import (
	"log"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitmq) Listen(exchangeName string, queueName string) (msgs <-chan amqp091.Delivery) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	if exchangeName != "" {
		err = ch.QueueBind(
			queue.Name,   // Nome da fila
			"",           // Routing Key (não usada em Fanout)
			exchangeName, // Nome do Exchange
			false,
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	msgs, err = ch.Consume(
		queue.Name,
		"",
		false, // Auto Ack: true (confirma automaticamente)
		false, // Exclusive: false
		false, // No Local: false
		false, // No Wait: false
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}
