package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something: ")
	consoleInput, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(consoleInput)

}
