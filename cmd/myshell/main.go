package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
			os.Exit(1)
		}

		args := strings.Fields(command)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			handleEcho(args)
		default:
			fmt.Printf("%s: command not found\n", args[0])
		}
	}
}

func handleEcho(args []string) {
	for i := 1; i < len(args); i++ {
		fmt.Print(args[i], " ")
	}
	fmt.Println()
}
