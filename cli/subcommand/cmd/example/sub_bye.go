package main

import (
	"flag"
	"fmt"
	"time"
)

func byeCmd() command {
	fset := flag.NewFlagSet(mainCmdName+" bye", flag.ExitOnError)
	opts := &byeOpts{date: time.Now()}
	fset.Var(&DateFlag{&opts.date}, "date", "Byebye date. (format: yyyymmdd)")

	return command{
		name:        "bye",
		description: "Print bye message with date",
		fset:        fset,
		fn: func(args []string, glOpts *globalOpts) error {
			fset.Parse(args)
			return bye(opts, glOpts)
		},
	}
}

type byeOpts struct {
	date time.Time
}

func bye(opts *byeOpts, glOpts *globalOpts) error {
	fmt.Printf("Bye-bye, subcommand on %s!\n", opts.date.Format("2006/01/02"))

	return nil
}
