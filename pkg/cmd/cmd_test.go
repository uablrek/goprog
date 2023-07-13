package cmd_test

import (
	"context"
	"testing"

	"github.com/uablrek/goprog/pkg/cmd"
)

func TestCmd(t *testing.T) {
	if cmd.Invoke(context.TODO(), nil) != 0 {
		t.Fatalf("Empty should work")
	}
	if cmd.Invoke(context.TODO(), []string{"nocmd"}) == 0 {
		t.Fatalf("Unknown command should fail")
	}
}
