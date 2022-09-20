package rabbitmq

import "github.com/streadway/amqp"

type Connect struct {
	Channel *amqp.Channel
}

func GetConnectRabbitMQ(rabbitUrl string) (Connect, error) {
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		return Connect{}, err
	}

	ch, err := conn.Channel()
	return Connect{Channel: ch}, err
}
