package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func weatherCmd() command {
	subCommands := []command{
		weatherGoodCmd(),
		weatherBadCmd(),
	}

	fset := flag.NewFlagSet(mainCmdName+" weather", flag.ExitOnError)
	fset.Usage = func() {
		fmt.Fprintln(fset.Output(), fmt.Sprintf("Usage: %s weather <sub-command> [command flags]", mainCmdName))
		fset.PrintDefaults()

		fmt.Fprintln(fset.Output())
		fmt.Fprintln(fset.Output(), "Sub Commands:")
		printCommands(fset, 0, subCommands)
	}
	opts := &weatherOpts{}

	return command{
		name:        "weather",
		description: "Print weather",
		Subcommands: subCommands,
		fset:        fset,
		fn: func(args []string, glOpts *globalOpts) error {
			fset.Parse(args)
			return weather(opts, glOpts)
		},
	}
}

type weatherOpts struct {
}

func weather(opts *weatherOpts, glOpts *globalOpts) error {
	return nil
}

func weatherGoodCmd() command {
	fset := flag.NewFlagSet(mainCmdName+" weather good", flag.ExitOnError)
	opts := &weatherGoodOpts{}

	return command{
		name:        "good",
		description: "Print weather is good",
		fset:        fset,
		fn: func(args []string, glOpts *globalOpts) error {
			fset.Parse(args)
			return weatherGood(opts, glOpts)
		},
	}
}

type weatherGoodOpts struct {
}

func weatherGood(opts *weatherGoodOpts, glOpts *globalOpts) error {
	// You can call core logic package's function.
	// fmt.Printf("The weather is %s!\n", forecast.GoodWeather())
	fmt.Printf("The weather is %s!\n", GoodWeather())
	return nil
}

func weatherBadCmd() command {
	fset := flag.NewFlagSet(mainCmdName+" weather bad", flag.ExitOnError)
	opts := &weatherBadOpts{}

	return command{
		name:        "bad",
		description: "Print weather is bad",
		fset:        fset,
		fn: func(args []string, glOpts *globalOpts) error {
			fset.Parse(args)
			return weatherBad(opts, glOpts)
		},
	}
}

type weatherBadOpts struct {
}

func weatherBad(opts *weatherBadOpts, glOpts *globalOpts) error {
	// You can call core logic package's function.
	// fmt.Printf("The weather is %s!\n", forecast.BadWeather())
	fmt.Printf("The weather is %s!\n", BadWeather())
	return nil
}

/*
  --------------- This code is just a copy of forecast package ---------------
  This is here because gonew doesn't support rewriting of package import path yet.
*/

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
