package subcommand

import (
	"context"
	"flag"
	"os"

	"github.com/go-logr/logr"
	"github.com/uablrek/goprog/pkg/cmd"
	flagsfiller "github.com/itzg/go-flagsfiller"
)

func init() {
	cmd.Cmds["subcommand"] = main
}

type config struct {
	Delay string `default:"1m30s" usage:"A delay in time.Duration format"`
	Help bool `usage:"Print help and exit" env:""`
}
var cfg config

func main(ctx context.Context, args []string) int {
	logger := logr.FromContextOrDiscard(ctx).WithValues("package", "subcommand")
	var flagset flag.FlagSet
	filler := flagsfiller.New(flagsfiller.WithEnv("TMPL"))
	if err := filler.Fill(&flagset, &cfg); err != nil {
		panic(err)
	}
	if err := flagset.Parse(args[1:]); err != nil {
		return 1
	}

	logger.V(1).Info("Called", "config", cfg)
	if cfg.Help {
		flagset.SetOutput(os.Stdout)
		flagset.PrintDefaults()
		return 0
	}

	return 0
}
