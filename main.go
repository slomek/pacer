package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	distance = flag.Float64("distance", 0.0, "total distance of the training")
	duration = flag.Duration("duration", 0, "total duration of the training")
	endo     = flag.Bool("endo", false, "print a comment for Endomondo")
)

func main() {
	flag.Parse()

	seconds := duration.Seconds()
	secPerKm := seconds / *distance

	roundedSecPerKmRoun := math.Ceil(secPerKm)
	pace, err := time.ParseDuration(fmt.Sprintf("%vs", roundedSecPerKmRoun))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pace (min/km): %v\n", pace)

	speed := math.Floor(100*3600.0/secPerKm) / 100.0
	fmt.Printf("Speed (km/h): %v\n", speed)

	if *endo {
		fmt.Printf("Distance: %vkm, ~%vmin/km, ~%vkm/h\n", *distance, pace, speed)
	}
}
