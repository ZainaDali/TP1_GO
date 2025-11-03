package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s!", name)
}
