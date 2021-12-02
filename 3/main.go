package main

import (
	"log"

	common "github.com/rbtr/aoc2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

func one(puzzle *common.Puzzle) error {
	return nil
}

func two(puzzle *common.Puzzle) error {
	return nil
}
