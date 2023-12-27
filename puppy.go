package puppy

import (
	"fmt"

	"github.com/AmanSarraf/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "woof woof woof !"

}
func BigBark() string {
	return dog.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowsUp(Barks())
}
func version() {
	fmt.Println("I am v1.0.0")
}
