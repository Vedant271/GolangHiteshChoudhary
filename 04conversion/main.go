package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Print("Please rate between 1 to 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Input rating is =", input)

	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if(error != nil){
		fmt.Println(error)
	} else {
		fmt.Println("Input rating after adding 1 =", numRating + 1)
	}
}

