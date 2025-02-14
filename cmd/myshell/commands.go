package main

import (
	"fmt"
	"os"
	"path/filepath"
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
			shellType: "a shell builtin",
			callback:  commandExit,
		},
		"echo": {
			name:      "echo",
			shellType: "a shell builtin",
			callback:  commandEcho,
		},
		"type": {
			name:      "type",
			shellType: "a shell builtin",
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

	cmd := args[1]
	c, ok := commandRegistry[cmd]
	if ok {
		fmt.Printf("%s is %s\n", c.name, c.shellType)
		return
	}

	pathVar := os.Getenv("PATH")
	paths := filepath.SplitList(pathVar)
	for _, path := range paths {
		f := filepath.Join(path, cmd)
		_, err := os.Stat(f)
		if err != nil {
			continue
		}
		fmt.Printf("%s is %s\n", cmd, f)
		return
	}

	fmt.Printf("%s: not found\n", cmd)
}
