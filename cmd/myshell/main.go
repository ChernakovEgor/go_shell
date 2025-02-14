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

		cm, ok := commandRegistry[args[0]]
		if ok {
			cm.callback(args...)
			continue
		}

		err = pathCommand(args...)
		if err != nil {
			fmt.Printf("%s: command not found\n", args[0])
		}
	}
}
