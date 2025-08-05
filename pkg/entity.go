package ouranosis

import (
	"context"
)

type Entity interface {
	Tick(context.Context, Tick) error
}
