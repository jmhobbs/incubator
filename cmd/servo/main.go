package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	pin := flag.String("pin", "PWM1", "PWM pin")
	flag.Parse()

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	p := gpioreg.ByName(*pin)
	if p == nil {
		log.Fatal("Failed to find PWM pin")
	}

	servo := device.NewServo(p)
	defer servo.Stop()

	servo.Rotate(0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter degree of rotation, 0-90")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		deg, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("Invalid input, %q\n", text)
		} else {
			servo.Rotate(deg)
		}
	}
}
