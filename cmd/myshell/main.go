package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
			os.Exit(1)
		}

		command = command[:len(command)-1]
		switch command {
		case "exit 0":
			handleExit()
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func handleExit() {
	os.Exit(0)
}
