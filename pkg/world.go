package ouranosis

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
)

type World struct {
	entities map[string]Entity
	space    map[string]Position

	TickRate time.Duration
}

func (w *World) Run(ctx context.Context) error {
	log := log.FromContext(ctx)
	log.Info("Running world", "rate", w.TickRate)
	t := time.NewTicker(w.TickRate)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case tick := <-t.C:
			if err := w.Tick(ctx, Tick(tick)); err != nil {
				return err
			}
		}
	}
}

func (w *World) Tick(ctx context.Context, t Tick) error {
	return nil
}
