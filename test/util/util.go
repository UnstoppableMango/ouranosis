package util

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/onsi/ginkgo/v2"
	"github.com/unmango/go/vcs/git"
)

var gitRoot string

func init() {
	gitRoot, _ = git.Root(context.Background())
}

func Run(cmd *exec.Cmd) (string, error) {
	ginkgo.GinkgoHelper()

	stderr, stdout := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Dir = gitRoot

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s: %w", stderr, err)
	} else {
		return stdout.String(), nil
	}
}
