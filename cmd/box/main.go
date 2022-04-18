package main

import (
	"context"
	"log"

	"github.com/kskitek/thebox/internal/pubsub"
	//"github.com/kskitek/thebox/internal/servo"
)

func main() {
	//servo := servo.New()
	//servo.CloseDoors(nil)

	sub, err := pubsub.NewSubscriber()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ch := make(chan pubsub.Message)
	go func() {
		err = sub.Subscribe(ctx, ch)
		if err != nil {
			panic(err)
		}
	}()

	for msg := range ch {
		log.Println(msg)
		switch msg {
		case pubsub.MessageOpen:
			//servo.OpenDoors(ctx)
		case pubsub.MessageClose:
			//servo.CloseDoors(ctx)
		}
	}
}
