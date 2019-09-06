package main

import (
	"flag"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	pin := flag.String("pin", "GPIO16", "pin trigger is on")
	flag.Parse()
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	p := gpioreg.ByName(*pin)
	if p == nil {
		log.Fatal("Failed to find pin")
	}

	p.Out(gpio.High)
	time.Sleep(1 * time.Second)
	p.Out(gpio.Low)
}
