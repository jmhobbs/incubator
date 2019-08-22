package main

import (
	"flag"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	var pin *string = flag.String("pin", "17", "GPIO pin the relay is connected to")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	relay := device.NewRelay(gpioreg.ByName(*pin))

	for {
		if relay.IsOpen() {
			color.Green("Relay is OPEN")
		} else {
			color.Red("Relay is CLOSED")
		}

		relay.Toggle()
		time.Sleep(time.Second * 2)
	}
}
