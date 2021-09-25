package servo

import (
	"context"
	"log"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func New() Servo {
	return Servo{}
}

type Servo struct {
}

func (s Servo) Open(ctx context.Context) {
	log.Println("openning")

	pin := "1"
	r := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(r, pin)

	err := servo.Start()
	if err != nil {
		panic(err)
	}

	err = servo.Max()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)

	err = servo.Min()
	if err != nil {
		panic(err)
	}

	err = servo.Center()
	if err != nil {
		panic(err)
	}
}

func (s Servo) Close(ctx context.Context) {
	log.Println("closing")
}
