package ouranosis_test

import (
	"context"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/unmango/go/vcs/git"
)

var (
	gitRoot    string
	clientPath string
	serverPath string
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	var err error
	gitRoot, err = git.Root(ctx)
	Expect(err).NotTo(HaveOccurred())

	clientPath, err = gexec.Build(filepath.Join(gitRoot, "cmd", "client"))
	Expect(err).NotTo(HaveOccurred())

	serverPath, err = gexec.Build(filepath.Join(gitRoot, "cmd", "server"))
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
