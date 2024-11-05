package kitten

import (
	"github.com/YanLOo97/cat"
)

func Meow() string {
	return "Meow!"
}

func Meows() string {
	return "Meow! Meow! Meow!"
}

func BigMeow() string {
	return cat.WhenGrownUp(Meow())
}

func BigMeows() string {
	return cat.WhenGrownUp(Meows())
}
