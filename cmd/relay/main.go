package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	var triggerPin *string = flag.String("trigger", "22", "GPIO pin the relay is connected to")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	relay := device.NewRelay(gpioreg.ByName(*triggerPin))

	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')
		relay.Toggle()

		if relay.IsClosed() {
			color.Green("Relay is CLOSED")
		} else {
			color.Red("Relay is OPEN")
		}
	}
}
