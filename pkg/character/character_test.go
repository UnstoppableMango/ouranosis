package character_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/ouranosis/pkg/character"
	"github.com/unstoppablemango/ouranosis/pkg/game"
)

var _ = Describe("Roll", func() {
	It("stays within the class range", func() {
		for sub := range uint64(100) {
			for _, c := range character.Classes {
				for _, s := range character.Stats {
					Expect(character.Roll(sub, c, s)).To(And(BeNumerically(">=", 4), BeNumerically("<=", 13)))
				}
			}
		}
	})

	It("favors the class stat", func() {
		Expect(character.Roll(0, character.Warrior, character.Strength)).To(BeNumerically(">", character.Roll(0, character.Mage, character.Strength)))
	})
})

var _ = Describe("Create", func() {
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC)

	It("rolls the same character for the same seed", func() {
		a, err := character.Create(ctx, game.Local{Seed: 1}, "Ana", character.Rogue, now)
		Expect(err).NotTo(HaveOccurred())
		b, err := character.Create(ctx, game.Local{Seed: 1}, "Ana", character.Rogue, now)
		Expect(err).NotTo(HaveOccurred())
		Expect(a).To(Equal(b))
		Expect(a.Progress).To(HaveLen(len(character.Stats)))
	})

	It("rolls differently for a different seed", func() {
		a, _ := character.Create(ctx, game.Local{Seed: 1}, "Ana", character.Rogue, now)
		b, _ := character.Create(ctx, game.Local{Seed: 2}, "Ana", character.Rogue, now)
		Expect(a.Base).NotTo(Equal(b.Base))
	})

	It("rejects an unknown class", func() {
		_, err := character.Create(ctx, game.Local{}, "Ana", "bard", now)
		Expect(err).To(HaveOccurred())
	})
})
