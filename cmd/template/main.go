package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	flagsfiller "github.com/itzg/go-flagsfiller"
	"github.com/uablrek/goprog/pkg/cmd"
	"github.com/uablrek/goprog/pkg/log"
	_ "github.com/uablrek/goprog/cmd/template/subcommand"
)

var version = "unknown"

var help = `
template [options] subcommand [options]
`

type config struct {
	LogLevel string `default:"info" usage:"debug|trace or an int"`
	Count    int    `default:"5" usage:"Some counter"`
	Yolo     bool   `default:"true" usage:"You Only Live Once"`
}

var cfg config

func main() {
	var flagset flag.FlagSet

	filler := flagsfiller.New(flagsfiller.WithEnv("TMPL"))
	if err := filler.Fill(&flagset, &cfg); err != nil {
		panic(err)
	}
	if err := flagset.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}
	ctx := log.NewContext(
		context.Background(), log.ZapLogger(cfg.LogLevel))
	logger := logr.FromContextOrDiscard(ctx)

	logger.V(1).Info("Start", "version", version, "config", cfg)
	cmd.Cmds["version"] = cmdVersion

	args := flagset.Args()
	if len(args) < 1 {
		fmt.Println(help)
		flagset.SetOutput(os.Stdout)
		flagset.PrintDefaults()
		cmd.PrintCommands()
		return
	}
	os.Exit(cmd.Invoke(ctx, args))
}

func cmdVersion(ctx context.Context, args []string) int {
	fmt.Println(version)
	return 0
}
