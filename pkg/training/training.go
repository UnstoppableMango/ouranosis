// Package training is the incremental sub-game that levels one stat.
package training

import (
	"context"
	"errors"
	"time"

	"github.com/unmango/game/curve"
	"github.com/unmango/game/num"

	"github.com/unstoppablemango/ouranosis/pkg/character"
	"github.com/unstoppablemango/ouranosis/pkg/game"
)

// Curve parameters for the first slice. Tuning lives here until it moves to
// a data file.
var (
	// UpgradeCost prices the nth upgrade.
	UpgradeCost = curve.Exponential{Base: num.FromInt(10), Growth: 1.15}
	// Rate is experience per second with n upgrades owned.
	Rate = curve.Exponential{Base: num.One(), Growth: 1.25}
	// LevelCost is the experience needed for the nth level above base.
	LevelCost = curve.Polynomial{Scale: num.FromInt(100), Degree: 1.5}
)

// ErrCannotAfford is returned by Buy when the balance is short.
var ErrCannotAfford = errors.New("training: cannot afford the next upgrade")

// Session trains one stat.
type Session struct {
	Calc     game.Calculator
	Stat     character.Stat
	Base     int64
	Progress *character.Progress
}

// Tick credits the experience produced since the last tick and returns it.
// The same call handles a one second tick and a week away.
func (s *Session) Tick(ctx context.Context, now time.Time) (num.Number, error) {
	elapsed := now.Sub(s.Progress.LastSeen)
	gained, err := s.Calc.Advance(ctx, Rate, s.Progress.Upgrades, elapsed)
	if err != nil {
		return num.Zero(), err
	}
	s.Progress.Earned = s.Progress.Earned.Add(gained)
	s.Progress.LastSeen = now
	return gained, nil
}

// RatePerSecond is the current experience rate.
func (s *Session) RatePerSecond(ctx context.Context) (num.Number, error) {
	return s.Calc.Evaluate(ctx, Rate, s.Progress.Upgrades)
}

// NextCost is the price of the next upgrade.
func (s *Session) NextCost(ctx context.Context) (num.Number, error) {
	return s.Calc.Evaluate(ctx, UpgradeCost, s.Progress.Upgrades)
}

// Affordable is how many upgrades the balance covers.
func (s *Session) Affordable(ctx context.Context) (int64, error) {
	return s.Calc.Invert(ctx, UpgradeCost, s.Progress.Upgrades, s.Progress.Balance())
}

// Buy purchases the next upgrade.
func (s *Session) Buy(ctx context.Context) error {
	cost, err := s.NextCost(ctx)
	if err != nil {
		return err
	}
	if s.Progress.Balance().Less(cost) {
		return ErrCannotAfford
	}
	s.Progress.Spent = s.Progress.Spent.Add(cost)
	s.Progress.Upgrades++
	return nil
}

// Level is the stat's level: base plus the levels the earned experience
// covers. Spending never lowers it.
func (s *Session) Level(ctx context.Context) (int64, error) {
	gained, err := s.Calc.Invert(ctx, LevelCost, 1, s.Progress.Earned)
	if err != nil {
		return 0, err
	}
	return s.Base + gained, nil
}
