package main

import (
	"fmt"
	"github.com/leoxiang66/Drone-Smart-Cities/environment"
)

func main() {
	for {
		flight := <-environment.RunningFlights

		fmt.Println(flight)

	}

}
