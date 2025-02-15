package main

import (
	"log"
	"os"
)

type State struct {
	CurrentDir string
}

func NewState() *State {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("getting current dir: %v", err)
	}
	return &State{
		CurrentDir: dir,
	}
}
