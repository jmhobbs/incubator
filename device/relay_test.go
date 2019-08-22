package device

import (
	"testing"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func TestNewRelay(t *testing.T) {
		pin := gpiotest.Pin{N: "in"}
		r := NewRelay(&pin)
		if pin.L != gpio.High {
			t.Error("Expected pin to be high, is low.")
		}
		if !r.IsOpen() {
			t.Error("Expected relay to be open, is closed.")
		}
		if r.IsClosed() {
			t.Error("Expected relay to be open, is closed.")
		}
}

func TestRelayOpenClose(t *testing.T) {
	pin := gpiotest.Pin{N: "in"}
	r := NewRelay(&pin)
	r.Close()
	if pin.L != gpio.Low {
		t.Error("Expected pin to be low, is high.")
	}
	if !r.IsClosed() {
		t.Error("Expected relay to be closed, is open.")
	}

	r.Open()
	if pin.L != gpio.High {
		t.Error("Expected pin to be high, is low.")
	}
	if !r.IsOpen() {
		t.Error("Expected relay to be open, is closed.")
	}
}

func TestRelayToggle(t *testing.T) {
	pin := gpiotest.Pin{N: "in"}
	r := NewRelay(&pin)

	r.Toggle()
	if pin.L != gpio.Low {
		t.Error("Expected pin to be low, is high.")
	}
	if !r.IsClosed() {
		t.Error("Expected relay to be closed, is open.")
	}

	r.Toggle()
	if pin.L != gpio.High {
		t.Error("Expected pin to be high, is low.")
	}
	if !r.IsOpen() {
		t.Error("Expected relay to be open, is closed.")
	}
}
