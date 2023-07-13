package cmd

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
)

var Cmds = make(map[string]func(ctx context.Context, args []string) int)

func Invoke(ctx context.Context, args []string) int {
	logger := logr.FromContextOrDiscard(ctx).WithValues("package", "cmd")
	if len(args) < 1 {
		fmt.Println("Subcommands:")
		for k := range Cmds {
			fmt.Println("  ", k)
		}
		return 0
	}

	if cmd, ok := Cmds[args[0]]; ok {
		logger.V(2).Info("Invoke", "args", args)
		rc := cmd(ctx, args)
		logger.V(2).Info("Return", "return-code", rc)
		return rc
	}
	logger.Info("Invalid", "command", args[0])
	return -1
}

func PrintCommands() {
	fmt.Println("\nSubcommands:")
	for k := range Cmds {
		fmt.Println("  ", k)
	}
	fmt.Println("")
}
