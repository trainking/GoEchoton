package simplemqx

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type SimpleQ struct {
	Name    string
	connurl string
	conn    *amqp.Connection
}

func New(connurl string, name string) *SimpleQ {
	return &SimpleQ{
		Name:    name,
		connurl: connurl,
	}
}

func (q *SimpleQ) Connect() error {
	conn, err := amqp.Dial(q.connurl)
	if err != nil {
		return err
	}
	q.conn = conn
	go func() {
		fmt.Printf("Closing: %s", <-conn.NotifyClose(make(chan *amqp.Error)))
		for err := q.Connect(); err != nil; err = q.Connect() {
			fmt.Println(err)
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

func (q *SimpleQ) Close() error {
	return q.conn.Close()
}

// Push 发送消息
func (q *SimpleQ) Push(body []byte) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	return err
}

// Cousume 消费消息
func (q *SimpleQ) Cousume(cFun func([]byte) error) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := cFun(d.Body); err == nil {
				d.Ack(false)
			}
		}
	}()
	return nil
}
