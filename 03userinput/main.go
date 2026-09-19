package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza: ")

	// comma ok || comma error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T\n", input)
}