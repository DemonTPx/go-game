package main

import (
	"fmt"

	"github.com/demontpx/go-game/lib/actor"
	"github.com/demontpx/go-game/lib/game"
)

func main() {
	loadActor()
	startGame()
}

func loadActor() {
	loader := actor.NewLoader()
	a, err := loader.LoadActorFromFile("res/actor/ball.yml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("actor loaded: %+v", a)
}

func startGame() {
	main := game.NewMain()
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
