package server_test

import (
	"context"

	gamev1alpha1 "buf.build/gen/go/unmango/game/protocolbuffers/go/dev/unmango/game/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ouranosisv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	"github.com/unstoppablemango/ouranosis/pkg/server"
)

var _ = Describe("Server", func() {
	var s *server.Server

	BeforeEach(func() {
		s = server.NewServer()
	})

	It("should reduce", func(ctx context.Context) {
		req := &ouranosisv1alpha1.ReduceRequest{
			State: &gamev1alpha1.State{},
		}

		res, err := s.Reduce(ctx, req)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.State).NotTo(BeNil())
		Expect(res.Events).NotTo(BeNil())
	})
})
