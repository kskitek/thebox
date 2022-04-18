package servo

import (
	"context"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	defaultPin = 19 // PWM1
)

func New() Servo {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	pwm := rpio.Pin(defaultPin)
	pwm.Mode(rpio.Pwm)
	rpio.SetFreq(pwm, 50*100)

	return Servo{
		pwm: pwm,
	}
}

type Servo struct {
	pwm rpio.Pin
}

func (s Servo) OpenDoors(ctx context.Context) {
	log.Println("openning")
	time.Sleep(time.Second / 2)
	s.pwm.DutyCycle(9, 100)

}

func (s Servo) CloseDoors(ctx context.Context) {
	log.Println("closing")
	time.Sleep(time.Second / 2)
	s.pwm.DutyCycle(3, 100)
}

func (s Servo) Close(ctx context.Context) {
	s.pwm.DutyCycle(3, 100)
	time.Sleep(time.Second / 2)
	s.pwm.Mode(rpio.Input)
	rpio.Close()
}
