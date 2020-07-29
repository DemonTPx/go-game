VERSION=0.1

.PHONY: all clean

all: build/go-game

clean:
	rm -rf build/*

build/go-game:
	go build -o build/go-game .
