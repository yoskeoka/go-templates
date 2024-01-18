package main

import (
	"flag"
	"fmt"

	"github.com/yoskeoka/go-tempaltes/cli/subcommand/forecast"
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
	fmt.Printf("The weather is %s!\n", forecast.GoodWeather())
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
	fmt.Printf("The weather is %s!\n", forecast.BadWeather())
	return nil
}
