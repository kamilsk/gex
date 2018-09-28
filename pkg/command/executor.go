package command

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Executor is an interface for executing commands.
type Executor interface {
	Exec(ctx context.Context, name string, args ...string) error
}

// NewExecutor creates a new Executor instance.
func NewExecutor(outW, errW io.Writer, inR io.Reader, cwd string) Executor {
	env := os.Environ()
	for _, e := range env {
		kv := strings.SplitN(e, "=", 2)
		if kv[0] == "PATH" {
			kv[1] = filepath.Join(cwd, "bin") + string(os.PathListSeparator) + kv[1]
		}
		env = append(env, strings.Join(kv, "="))
	}
	return &executorImpl{
		outW: outW,
		errW: errW,
		inR:  inR,
		cwd:  cwd,
		env:  env,
	}
}

type executorImpl struct {
	outW, errW io.Writer
	inR        io.Reader
	cwd        string
	env        []string
}

func (e *executorImpl) Exec(ctx context.Context, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = e.outW
	cmd.Stderr = e.errW
	cmd.Stdin = e.inR
	cmd.Dir = e.cwd
	cmd.Env = e.env
	return cmd.Run()
}