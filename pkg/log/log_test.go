package log_test

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/uablrek/goprog/pkg/log"
)

func TestLog(t *testing.T) {
	ctx := log.NewContext(context.TODO(), log.ZapLogger("trace"))
	logger := logr.FromContextOrDiscard(ctx)
	logger.Info("An Info level entry", "number", 5)
	logger.V(1).Info("A Debug level entry", "number", 5)
	logger.V(2).Info("A Trace level entry", "number", 5)
	logger.V(2).V(1).Info("This doesn't work. Should it?")

	if !logger.V(2).Enabled() {
		t.Fatal("Trace level not enabled")
	}
	if l := logger.V(2); l.Enabled() {
		l.Info("Some hard-to-get param", "value", "HARD")
	}

	ctx = log.NewContext(ctx, log.ZapLogger("debug")) // should be a no-op
	logger = logr.FromContextOrDiscard(ctx)
	logger.V(2).Info("Trace level still seen", "number", 5)

	ctx = log.NewContext(context.TODO(), log.ZapLogger(""))
	logger = logr.FromContextOrDiscard(ctx)
	logger.Info("Info level still seen", "number", 5)
	logger.V(1).Info("Debug suspressed", "number", 5)
	logger.V(1).Info("Trace suspressed", "number", 5)

	ctx = log.NewContext(context.TODO(), log.ZapLogger("10"))
	logger = logr.FromContextOrDiscard(ctx)
	logger.V(10).Info("Never used level")
}
