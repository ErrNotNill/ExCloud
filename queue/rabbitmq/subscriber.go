package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func (conn Connect) StartSubscribe(queueName, routingKey string, handler func(d amqp.Delivery) bool, concurrency int) error {
	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	err = conn.Channel.QueueBind(queueName, routingKey, "events", false, nil)
	if err != nil {
		return err
	}
	prefetchCount := concurrency * 4
	err = conn.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
	}
	msgs, err := conn.Channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for i := 0; i < concurrency; i++ {
		fmt.Printf("Processing messages on thread %v...\n", i)
		go func() {
			for msg := range msgs {
				if handler(msg) {
					msg.Ack(false)
				} else {
					msg.Nack(false, true)
				}
			}
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
		}()
	}
	return nil
}
