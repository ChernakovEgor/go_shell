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
		command = command[:len(command)-1]
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s: command not found\n", command)
	}
}
