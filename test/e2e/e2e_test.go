package ouranosis_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("E2E Suite", func() {
	Context("Server", func() {
		Describe("Lifecycle", Ordered, func() {
			var serverSession *gexec.Session

			It("should run", func() {
				cmd := exec.Command(serverPath)
				ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(ses).Should(gbytes.Say(`.*`))
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

				Eventually(ses).Should(gbytes.Say(`.*`))
				clientSession = ses
			})

			It("should exit", func() {
				Eventually(clientSession.Interrupt()).Should(gexec.Exit())
			})
		})
	})
})
