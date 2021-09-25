package main

import (
	"context"
	"log"

	"github.com/kskitek/thebox/internal/pubsub"
)

func main() {
	// servo := servo.New()
	sub, err := pubsub.NewSubscriber()
	if err != nil {
		panic(err)
	}

	ch := make(chan pubsub.Message)
	go func() {
		err = sub.Subscribe(context.Background(), ch)
		if err != nil {
			panic(err)
		}
	}()

	for msg := range ch {
		log.Println(msg)
	}
}
