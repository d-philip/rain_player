package main

import (
	"machine"
	"time"
)

const (
	led    = machine.Pin(13)
	button = machine.Pin(12)
)

func main() {
	// led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
