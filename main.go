package main

import (
	"runtime"

	"github.com/DemonTPx/go-game/lib/game"
)

func main() {
	runtime.LockOSThread()

	main := game.NewMain()
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
