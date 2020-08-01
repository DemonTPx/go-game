package main

import (
	"fmt"
	"runtime"

	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/game"
)

func main() {
	runtime.LockOSThread()

	loadActor()
	startGame()
}

func loadActor() {
	loader := actor.NewLoader()
	a, err := loader.LoadActorFromFile("res/actor/ball.yml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("actor loaded: %+v\n", a)
}

func startGame() {
	main := game.NewMain()
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
