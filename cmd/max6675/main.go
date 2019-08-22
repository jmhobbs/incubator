package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func main() {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	p, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	sensor, err := device.NewMAX6675(p)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := sensor.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%0.2f °C\n", c)
		time.Sleep(time.Second * 2)
	}
}
