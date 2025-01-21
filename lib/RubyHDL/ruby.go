package RubyHDL

import (
	"context"
	"log"
	"os/exec"
	"time"
)

type RubyHDL struct {
	rc      string
	re      string
	timeout time.Duration
}

func NewRubyHDL(rc, re string, timeout time.Duration) *RubyHDL {
	return &RubyHDL{
		rc:      rc,
		re:      re,
		timeout: timeout,
	}
}

func (r *RubyHDL) Re(dir string, rbs string, input string) (string, error) {
	ctx := context.Background()
	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, r.re, "-r", rbs, input)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	log.Println("[RE]", dir, string(out), err)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r *RubyHDL) Rc(dir string, rby string) (string, error) {
	ctx := context.Background()
	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}
	cmd := exec.CommandContext(ctx, "sml", "@SMLload="+r.rc, rby)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	log.Println("[RC]", dir, string(out), err, cmd, cmd.Args)
	return string(out), err
}
