package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type CmdOptions struct {
	User   string
	Host   string
	Port   int
	Key    string // path to private key (optional)
	DryRun bool
}

func RunSSH(opts CmdOptions) error {
	args := []string{}
	if opts.Key != "" {
		args = append(args, "-i", opts.Key)
	}
	if opts.Port > 0 && opts.Port != 22 {
		args = append(args, "-p", strconv.Itoa(opts.Port))
	}
	args = append(args, fmt.Sprintf("%s@%s", opts.User, opts.Host))

	if opts.DryRun {
		fmt.Println("ssh", quoteArgs(args))
		return nil
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func quoteArgs(args []string) string {
	out := ""
	for _, a := range args {
		if len(out) > 0 {
			out += " "
		}
		out += fmt.Sprintf("%q", a)
	}
	return out
}
