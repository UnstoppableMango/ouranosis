package training_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/game/num"

	"github.com/unstoppablemango/ouranosis/pkg/character"
	"github.com/unstoppablemango/ouranosis/pkg/game"
	"github.com/unstoppablemango/ouranosis/pkg/training"
)

var _ = Describe("Session", func() {
	var (
		ctx   = context.Background()
		start = time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC)
		sess  *training.Session
	)

	BeforeEach(func() {
		sess = &training.Session{
			Calc:     game.Local{Seed: 42},
			Stat:     character.Strength,
			Base:     7,
			Progress: &character.Progress{LastSeen: start},
		}
	})

	It("credits experience for elapsed time", func() {
		gained, err := sess.Tick(ctx, start.Add(10*time.Second))
		Expect(err).NotTo(HaveOccurred())
		Expect(gained.Float64()).To(BeNumerically("~", 10, 1e-9))
		Expect(sess.Progress.Balance().Float64()).To(BeNumerically("~", 10, 1e-9))
		Expect(sess.Progress.LastSeen).To(Equal(start.Add(10 * time.Second)))
	})

	It("credits a week away in one call", func() {
		gained, err := sess.Tick(ctx, start.Add(7*24*time.Hour))
		Expect(err).NotTo(HaveOccurred())
		Expect(gained.Float64()).To(BeNumerically("~", 604800, 1e-6))
	})

	It("refuses an upgrade the balance cannot cover", func() {
		Expect(sess.Buy(ctx)).To(MatchError(training.ErrCannotAfford))
		Expect(sess.Progress.Upgrades).To(BeZero())
	})

	It("buys an upgrade and raises the rate", func() {
		_, err := sess.Tick(ctx, start.Add(time.Minute))
		Expect(err).NotTo(HaveOccurred())
		before, _ := sess.RatePerSecond(ctx)
		Expect(sess.Buy(ctx)).To(Succeed())
		after, _ := sess.RatePerSecond(ctx)
		Expect(sess.Progress.Upgrades).To(Equal(int64(1)))
		Expect(sess.Progress.Spent.Float64()).To(BeNumerically("~", 10, 1e-9))
		Expect(after.Float64()).To(BeNumerically(">", before.Float64()))
	})

	It("counts affordable upgrades", func() {
		sess.Progress.Earned = num.FromInt(40)
		Expect(sess.Affordable(ctx)).To(Equal(int64(3)))
	})

	It("levels from earned experience and never from spent", func() {
		lvl, err := sess.Level(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(lvl).To(Equal(int64(7)))

		sess.Progress.Earned = num.FromInt(400)
		lvl, _ = sess.Level(ctx)
		Expect(lvl).To(Equal(int64(9)))

		sess.Progress.Spent = num.FromInt(399)
		after, _ := sess.Level(ctx)
		Expect(after).To(Equal(lvl))
	})
})
