package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name:")
	input ,err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Found an error: %s\n", err)
	} else {
		input = input[:len(input)-1]
		fmt.Printf("Hello, %s!\n", input)
	}
}
