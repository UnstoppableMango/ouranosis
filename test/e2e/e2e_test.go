package ouranosis_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/unstoppablemango/ouranosis/test/util"
)

var _ = Describe("E2E Suite", func() {
	It("Should connect", func() {
		server := util.Start(serverPath)
		client := util.Start(clientPath, "text")

		Eventually(client.Err).Should(gbytes.Say(`Got response`))
		Eventually(client).Should(gexec.Exit(0))
		Eventually(server.Interrupt()).Should(gexec.Exit())
	})
})
