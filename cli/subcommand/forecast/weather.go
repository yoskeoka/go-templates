package forecast

// This file is an example of command line's main core logic.

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// GoodWeather returns a good weather.
func GoodWeather() string {
	candidates := []string{
		"Sunny", "Cloudy",
	}
	return candidates[rand.Intn(len(candidates))]
}

// BadWeather returns a bad weather.
func BadWeather() string {
	candidates := []string{
		"Rainy", "Snowy", "Windy", "Foggy", "Thunderstorm",
	}
	return candidates[rand.Intn(len(candidates))]
}
