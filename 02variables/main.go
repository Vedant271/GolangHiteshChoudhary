package main

import "fmt"

const LoginToken string = "jfasfiwij"

func main() {
	var username string = "Hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)

	var smallFloat float32 = 2.234567823456782345678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 2.234567823456782345678
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// implicit types
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
