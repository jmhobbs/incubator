package main

import (
	"flag"
	"log"

	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	bus := flag.String("bus", "I2C1", "i2c bus name (like 'I2C0')")
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open(*bus)
	if err != nil {
		log.Fatal("error opening i2c bus:", err)
	}
	defer b.Close()

	dev := device.NewBME280(b)
	temperature, pressure, humidity, err := dev.Read()
	if err != nil {
		log.Fatal("error reading device:", err)
	}
	log.Printf("%0.2f C", temperature)
	log.Printf("%0.2f hPa", pressure)
	log.Printf("%0.2f%%", humidity)
}
