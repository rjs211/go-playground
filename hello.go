package main

import (
	"fmt"
	"time"

	"github.com/mshafiee/swephgo"
)

func getHouseCusps(utcTime string, latitude, longitude float64, houseSystem string) ([]float64, error) {
	// Parse UTC time
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing UTC time: %w", err)
	}

	var dret [2]float64
	var serr [256]byte

	// Convert UTC time to Julian Day number
	err_code := swephgo.UtcToJd(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), float64(t.Second()), swephgo.SeGregCal, dret[:], serr[:])
	if err_code == swephgo.Err {
		return nil, fmt.Errorf("error converting UTC time to Julian Day: %w", serr)
	}

	tjdUt := dret[1]

	// Calculate house cusps
	var cusps [13]float64
	var ascmc [10]float64
	swephgo.Houses(tjdUt, latitude, longitude, int('P'), cusps[:], ascmc[:])
	// if err_code == swephgo.Err {
	// 	return nil, fmt.Errorf("error converting UTC time to Julian Day: %w", serr)
	// }

	return cusps[:], nil
}

func main() {
	utcTime := "2024-02-27T00:00:00Z"
	latitude := 37.7833
	longitude := -122.4167

	cusps, err := getHouseCusps(utcTime, latitude, longitude, "P") // Using Placidus house system for example
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("House cusps:")
	for i, cusp := range cusps {
		fmt.Printf("House %d: %.6f degrees\n", i+1, cusp)
	}
}
