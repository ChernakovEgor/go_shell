package main

import (
	"fmt"
	"os"
)

type command struct {
	name      string
	shellType string
	callback  func(args ...string)
}

var commandRegistry map[string]command

func init() {
	commandRegistry = map[string]command{
		"exit": {
			name:      "exit",
			shellType: "a shell built in",
			callback:  commandExit,
		},
		"echo": {
			name:      "echo",
			shellType: "a shell built in",
			callback:  commandEcho,
		},
		"type": {
			name:      "type",
			shellType: "a shell built in",
			callback:  commandType,
		},
	}
}

func commandExit(args ...string) {
	os.Exit(0)
}

func commandEcho(args ...string) {
	for i := 1; i < len(args); i++ {
		fmt.Print(args[i], " ")
	}
	fmt.Println()
}

func commandType(args ...string) {
	if len(args) == 1 {
		return
	}
	c, ok := commandRegistry[args[1]]
	if !ok {
		fmt.Printf("%s: not found\n", args[1])
		return
	}
	fmt.Printf("%s is %s\n", c.name, c.shellType)
}
