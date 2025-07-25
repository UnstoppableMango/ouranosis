package main

import "math/rand/v2"

type Event struct{}

type Entity struct {
	X, Y int
}

type Brain struct {
}

func (b Brain) Apply(e Entity, _ []Event) (Entity, []Event) {
	if rand.Float32() < 0.3 {
		e.X++
	}

	return e, nil
}

func main() {

}
