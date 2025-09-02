package domain

type Entity interface {
	At(x, y int) bool
}

type Player interface {
	Entity

	Up()
	Down()
	Left()
	Right()
}
