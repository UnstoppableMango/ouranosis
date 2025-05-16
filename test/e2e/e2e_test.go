package ouranosis_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/unstoppablemango/ouranosis/test/util"
)

var _ = Describe("E2E Suite", func() {
	It("Should connect", func() {
		server := util.Start(serverPath)
		client := util.Start(clientPath)

		Eventually(client.Err).Should(gbytes.Say(`Got response`))
		Eventually(client).Should(gexec.Exit(0))
		Eventually(server.Interrupt()).Should(gexec.Exit())
	})

	Context("Server", func() {
		Describe("Lifecycle", Ordered, func() {
			var serverSession *gexec.Session

			It("should run", func() {
				cmd := exec.Command(serverPath)
				ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(ses.Err).Should(gbytes.Say(`ListenAndServing`))
				serverSession = ses
			})

			It("should exit", func() {
				Eventually(serverSession.Interrupt()).Should(gexec.Exit())
			})
		})
	})

	Context("Client", func() {
		Describe("Lifecycle", Ordered, func() {
			var clientSession *gexec.Session

			It("should run", func() {
				cmd := exec.Command(clientPath)
				ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(ses.Err).Should(gbytes.Say(`Got response`))
				clientSession = ses
			})

			It("should exit", func() {
				Eventually(clientSession).Should(gexec.Exit(0))
			})
		})
	})
})
