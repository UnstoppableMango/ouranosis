package ouranosis_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/go/vcs/git"
)

var gitRoot string

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	var err error
	gitRoot, err = git.Root(ctx)
	Expect(err).NotTo(HaveOccurred())
})
