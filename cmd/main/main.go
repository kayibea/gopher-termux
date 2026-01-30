package main

import (
	"fmt"
	"log"

	"github.com/kayibea/gopher-termux/battery"
)

func main() {
	s, err := battery.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Battery: %d%% (%s)\n", s.Percentage, s.Status)
}
