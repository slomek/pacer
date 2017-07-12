package main

import (
	"flag"
	"fmt"
	"math"
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

	pace := formatPace(secPerKm)
	fmt.Printf("Pace (min/km): %v\n", pace)

	speed := math.Floor(100*3600.0/secPerKm) / 100.0
	fmt.Printf("Speed (km/h): %v\n", speed)

	if *endo {
		fmt.Printf("Distance: %vkm, ~%vmin/km, ~%vkm/h\n", *distance, pace, speed)
	}
}

func formatPace(secsPerKm float64) string {
	secs := int(secsPerKm)
	mins, secs := int(math.Floor(float64(secs/60))), secs%60
	return fmt.Sprintf("%02d:%02d", int(mins), int(secs))
}
