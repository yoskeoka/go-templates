package main

import (
	"flag"
	"fmt"
	"strings"
)

func helloCmd() command {
	fset := flag.NewFlagSet(mainCmdName+" hello", flag.ExitOnError)
	opts := &helloOpts{}
	fset.Func("msg", "Multiple messages after Hello.", func(v string) error {
		opts.msg = append(opts.msg, v)
		return nil
	})

	return command{
		name:        "hello",
		description: "Print hello message",
		fset:        fset,
		fn: func(args []string, glOpts *globalOpts) error {
			fset.Parse(args)
			return hello(opts, glOpts)
		},
	}
}

type helloOpts struct {
	msg []string
}

func hello(opts *helloOpts, glOpts *globalOpts) error {
	fmt.Printf("Hello, subcommand! %s\n", strings.Join(opts.msg, ", "))

	return nil
}
