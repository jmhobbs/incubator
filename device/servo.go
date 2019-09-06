package device

import (
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/physic"
)

type Servo struct {
	pin gpio.PinIO
}

func NewServo(pin gpio.PinIO) *Servo {
	return &Servo{pin}
}

func (s *Servo) Stop() {
	s.pin.Halt()
}

func (s *Servo) Rotate(deg int) error {
	if deg > 90 {
		deg = 90
	}
	divisor := int(8 + (0.22 * float64(deg)))
	return s.pin.PWM(gpio.DutyMax/gpio.Duty(divisor), 50*physic.Hertz)
}
