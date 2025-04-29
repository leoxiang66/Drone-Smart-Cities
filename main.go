package main

import (
	"github.com/leoxiang66/Drone-Smart-Cities/environment"
	"fmt"
)


func main()  {
	districts, time := environment.SampleFlight()
	fmt.Println(districts)
	fmt.Println(time)
}