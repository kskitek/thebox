package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
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

	time.Sleep(time.Second * 2)

	err = servo.Center()
	if err != nil {
		panic(err)
	}
}
