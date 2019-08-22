package device

import (
	"periph.io/x/periph/conn/gpio"
)

// Considered to be a normally-open relay, triggered by a low pin.
type Relay struct {
	pin  gpio.PinIO
	open bool
}

func NewRelay(pin gpio.PinIO) *Relay {
	pin.Out(gpio.High) // Default open
	return &Relay{pin, true}
}

func (r *Relay) Open() {
	if !r.open {
		r.pin.Out(gpio.High)
		r.open = true
	}
}

func (r *Relay) Close() {
	if r.open {
		r.pin.Out(gpio.Low)
		r.open = false
	}
}

func (r *Relay) Toggle() {
	if r.open {
		r.Close()
	} else {
		r.Open()
	}
}

func (r *Relay) IsOpen() bool {
	return r.open
}

func (r *Relay) IsClosed() bool {
	return !r.open
}
