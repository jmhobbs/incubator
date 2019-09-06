package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/host"
)

func main() {
	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("== Loaded ==")
	for _, d := range state.Loaded {
		fmt.Println(d.String())
	}

	fmt.Println("\n== Skipped ==")
	for _, d := range state.Skipped {
		fmt.Println(d.String())
	}

	fmt.Println("\n== Failed ==")
	for _, d := range state.Failed {
		fmt.Println(d.String())
	}
}
