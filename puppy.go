package puppy

import (
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
