package kitten

import (
	"fmt"

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

func From11() {
	fmt.Println("I'm from version 1.1.0")
}