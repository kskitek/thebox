package pubsub

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func NewPublisher() (*Publisher, error) {
	client, err := pubsub.NewClient(context.Background(), os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, err
	}

	topic := os.Getenv("TOPIC")
	if topic == "" {
		return nil, errors.New("topic cannot be empty")
	}

	return &Publisher{
		client: client,
		topic:  topic,
	}, nil
}

type Publisher struct {
	client *pubsub.Client
	topic  string
}

func (p Publisher) Publish(ctx context.Context, message Message) error {
	topic := p.client.Topic(p.topic)
	msg := &pubsub.Message{Data: []byte(message)}
	_, err := topic.Publish(ctx, msg).Get(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewSubscriber() (*Subscriber, error) {
	client, err := pubsub.NewClient(context.Background(), os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, err
	}

	topic := os.Getenv("TOPIC")
	if topic == "" {
		return nil, errors.New("topic cannot be empty")
	}

	return &Subscriber{
		client: client,
		topic:  topic,
	}, nil
}

type Subscriber struct {
	client *pubsub.Client
	topic  string
}

func (s Subscriber) Subscribe(ctx context.Context, receiveChan chan<- Message) error {
	log.Println("subscribing to " + s.topic)
	sub := s.client.Subscription("thebox-sub")
	err := sub.Receive(ctx, s.handleMessage(receiveChan))
	if err != nil {
		return err
	}
	return nil
}

func (s Subscriber) handleMessage(ch chan<- Message) func(context.Context, *pubsub.Message) {
	return func(ctx context.Context, m *pubsub.Message) {
		ch <- Message(m.Data)
		m.Ack()
	}
}
