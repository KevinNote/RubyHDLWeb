package RubyDHL

import (
	"log"
	"os/exec"
)

type RubyDHL struct {
	rc string
	re string
}

func NewRubyDHL(rc, re string) *RubyDHL {
	return &RubyDHL{
		rc: rc,
		re: re,
	}
}

func (r *RubyDHL) Re(dir string, rbs string, input string) (string, error) {
	cmd := exec.Command(r.re, "-r", rbs, input)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	log.Println("[RE]", dir, string(out), err)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r *RubyDHL) Rc(dir string, rby string) (string, error) {
	cmd := exec.Command("sml", "@SMLload="+r.rc, rby)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	log.Println("[RC]", dir, string(out), err, cmd, cmd.Args)
	return string(out), err
}
