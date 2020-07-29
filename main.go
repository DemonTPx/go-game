package main

import (
	"github.com/demontpx/go-game/lib/game"
)

func main() {
	main := game.NewMain()
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
