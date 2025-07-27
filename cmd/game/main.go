package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Event struct{}

type Entity struct {
	X, Y int
}

type Brain struct {
}

func (b Brain) Apply(e Entity, _ []Event) (Entity, []Event) {
	if rand.Float32() < 0.3 {
		if rand.Float32() > 0.5 {
			e.X++
		} else {
			e.X--
		}
	}

	if rand.Float32() < 0.3 {
		if rand.Float32() > 0.5 {
			e.Y++
		} else {
			e.Y--
		}
	}

	return e, nil
}

func main() {
	e := Entity{}
	b := Brain{}

	for {
		time.Sleep(1 * time.Second)
		e, _ = b.Apply(e, nil)
		fmt.Printf("x: %d, y: %d\n", e.X, e.Y)
	}
}
