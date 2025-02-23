package rabbitmq

import (
	"log"
	"os"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type rabbitmq struct {
	Conn *amqp091.Connection
}

type RabbitMQInterface interface {
	CloseConn()
	Publish(exchangeName string, exchangeType string, data []byte)
	Listen(exchangeName string, queueName string) (msgs <-chan amqp091.Delivery)
}

func NewRabbitMQConnection() RabbitMQInterface {
	conn, err := amqp091.Dial(os.Getenv("RABBIT_MQ"))
	if err != nil {
		log.Fatal(err)
	}

	return &rabbitmq{
		Conn: conn,
	}
}

func (r *rabbitmq) CloseConn() {
	r.Conn.Close()
}
