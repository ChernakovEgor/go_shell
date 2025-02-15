package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	// "os/exec"
	"path/filepath"
)

type command struct {
	name      string
	shellType string
	callback  func(*State, ...string)
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
		"pwd": {
			name:      "pwd",
			shellType: "a shell builtin",
			callback:  commandPwd,
		},
		"cd": {
			name:      "cd",
			shellType: "a shell builtin",
			callback:  commandCd,
		},
	}
}

func commandExit(_ *State, args ...string) {
	os.Exit(0)
}

func commandEcho(_ *State, args ...string) {
	for i := 1; i < len(args); i++ {
		fmt.Print(args[i], " ")
	}
	fmt.Println()
}

func commandType(_ *State, args ...string) {
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

func pathCommand(_ *State, args ...string) error {
	cmnd := args[0]

	pathVar := os.Getenv("PATH")
	paths := filepath.SplitList(pathVar)
	for _, path := range paths {
		f := filepath.Join(path, cmnd)
		stat, err := os.Stat(f)
		if err != nil || stat.Mode().Perm()&0111 != 0111 {
			continue
		}

		cmd := exec.Command(cmnd, args[1:]...)
		res, err := cmd.Output()
		fmt.Print(string(res))
		return nil
	}

	return fmt.Errorf("%s not in PATH", cmnd)
}

func commandPwd(state *State, args ...string) {
	fmt.Println(state.CurrentDir)
}

func commandCd(state *State, args ...string) {
	if len(args) == 1 {
		fmt.Println("go to home")
		return
	}

	p := args[1]

	if filepath.IsAbs(p) {
		_, err := os.Stat(p)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", p)
			return
		}
	} else if p[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("getting home dir: %v", err)
		}
		p = home
	} else {
		p = filepath.Join(state.CurrentDir, p)
	}

	err := os.Chdir(p)
	if err != nil {
		fmt.Printf("cd: %s: could not change directory\n", p)
		return
	}

	state.CurrentDir = p
}
