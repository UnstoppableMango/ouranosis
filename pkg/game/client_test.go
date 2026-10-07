package game_test

import (
	"context"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/game/curve"
	"github.com/unmango/game/identity"
	"github.com/unmango/game/num"
	"github.com/unmango/game/server"

	"github.com/unstoppablemango/ouranosis/pkg/game"
)

// The client against a real framework handler must agree with Local for the
// same seed, since both are the same pure functions.
var _ = Describe("Client", func() {
	var (
		ctx    = context.Background()
		client *game.Client
		local  game.Local
		exp    = curve.Exponential{Base: num.FromInt(10), Growth: 1.15}
	)

	BeforeEach(func() {
		id, err := identity.Generate(time.Now())
		Expect(err).NotTo(HaveOccurred())
		ts := httptest.NewServer(server.Handler(id))
		DeferCleanup(ts.Close)
		client = game.New(ts.URL)
		local = game.Local{Seed: id.Seed}
	})

	It("derives", func() {
		want, _ := local.Derive(ctx, "character/strength")
		Expect(client.Derive(ctx, "character/strength")).To(Equal(want))
	})

	It("evaluates", func() {
		want, _ := local.Evaluate(ctx, exp, 10)
		Expect(client.Evaluate(ctx, exp, 10)).To(Equal(want))
	})

	It("inverts", func() {
		want, _ := local.Invert(ctx, exp, 0, num.FromInt(40))
		Expect(client.Invert(ctx, exp, 0, num.FromInt(40))).To(Equal(want))
	})

	It("advances", func() {
		want, _ := local.Advance(ctx, exp, 3, time.Minute)
		Expect(client.Advance(ctx, exp, 3, time.Minute)).To(Equal(want))
	})
})
