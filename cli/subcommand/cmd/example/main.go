package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	exitCode := cliMain()
	os.Exit(exitCode)
}

var (
	mainCmdName = "example"
)

type globalOpts struct {
	output io.Writer
}

type command struct {
	name        string
	description string
	Subcommands []command
	fset        *flag.FlagSet
	fn          func(args []string, gOpts *globalOpts) error
}

func cliMain() int {
	commands := []command{
		helloCmd(),
		byeCmd(),
		weatherCmd(),
	}

	fset := flag.NewFlagSet(mainCmdName, flag.ExitOnError)
	version := fset.Bool("version", false, "Print version")

	glOpts := &globalOpts{
		output: os.Stdout,
	}

	fset.Usage = func() {
		fmt.Fprintln(fset.Output(), fmt.Sprintf("Usage: %s <command> [command flags]", mainCmdName))
		fset.PrintDefaults()

		fmt.Fprintln(fset.Output())
		fmt.Fprintln(fset.Output(), "Commands:")
		printCommands(fset, 0, commands)
	}

	fset.Parse(os.Args[1:])

	if *version {
		fmt.Fprintf(fset.Output(), "Version: %s\n", CommitHash)
		return 0
	}

	args := fset.Args()
	if len(args) == 0 {
		fset.Usage()
		return 1
	}

	err := runSubcmd(mainCmdName, commands, args, glOpts)
	if err != nil {
		fmt.Fprint(fset.Output(), err)
		return 1
	}

	return 0
}

func printCommands(fset *flag.FlagSet, indent int, commands []command) {
	for _, cmd := range commands {
		if cmd.fset == nil || cmd.fn == nil {
			continue // skip not implemented
		}

		fmt.Fprintf(fset.Output(), "%s  %s:%s%s\n", strings.Repeat(" ", indent), cmd.name, strings.Repeat(" ", 12-len(cmd.name)), cmd.description)
	}
}

func runSubcmd(parentCmd string, subCommands []command, args []string, glOpts *globalOpts) error {

	subCmd := args[0]
	for _, cmd := range subCommands {
		if cmd.name == subCmd {

			if len(cmd.Subcommands) > 0 {
				if len(args) == 1 {
					cmd.fset.Usage()
					return nil
				}

				err := runSubcmd(cmd.name, cmd.Subcommands, args[1:], glOpts)
				if err != nil {
					return err
				}
			}

			err := cmd.fn(args[1:], glOpts)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("unknown command: '%s' for '%v'\n", subCmd, parentCmd)
}
