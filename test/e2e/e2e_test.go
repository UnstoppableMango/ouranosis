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
	Context("Server", func() {
		Describe("Lifecycle", Ordered, func() {
			var (
				serverPath    string
				serverSession *gexec.Session
			)

			It("should build", func() {
				cmd := exec.Command("make", "bin/server")
				_, err := util.Run(cmd)
				Expect(err).NotTo(HaveOccurred())
				serverPath = gitRoot + "/bin/server"
				Expect(serverPath).To(BeARegularFile())
			})

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
})
