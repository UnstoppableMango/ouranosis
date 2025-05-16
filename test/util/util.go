package util

import (
	"context"
	"os/exec"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/unmango/go/vcs/git"
)

var gitRoot string

func init() {
	gitRoot, _ = git.Root(context.Background())
}

func Start(command string, arg ...string) *gexec.Session {
	ginkgo.GinkgoHelper()

	cmd := exec.Command(command, arg...)
	ses, err := gexec.Start(cmd, ginkgo.GinkgoWriter, ginkgo.GinkgoWriter)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	return ses
}
