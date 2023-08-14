package main

import "fmt"


// This will Throw an error --> outside any method this type of declaration is not allowed . It is only allowed inside methods or functions.
// jwtToken := 300000 

// but you can declare using var and dataType . It will correctly work. 
var jwtToken int = 300000


// Declaring Constant
// NOTE : since we have made the starting of the variable as capital letter ==> This signifies that the variable is PUBLIC Variable. 
const LoginToken string = "fsgdgsdgs" // Constants are not going to change . PUBLIC variable -> So this variable can be accessed from anywhere. 





func main() {
	// Creating variables 

	// String 
	var username string = "poorvaditya"
	fmt.Println(username)

	fmt.Printf("Variable is of type: %T\n", username) // %T is used to check the datatype of the variable. so here username is of type string.

	// Boolean 
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	// uint8 variablem -- Read DOCs
	var smallVal uint8 = 255 // it have maxRange of 255 . 256 -> will show error . Overflow.
    fmt.Println(smallVal)
    fmt.Printf("Variable is of type: %T\n", smallVal)

	// uint , uint64, uint32, -> all these are helpful when you are doing some OS work. 
	// Similarly for float (float32, float64)


	// DeFault values and some aliases 

	// declaring a variable 
	var anotherVariable int ;
    fmt.Println(anotherVariable) // Note : this will not put any garbage value in when we have not initialized a variable. By Defaul - it will have 0 value. 
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// Implicit type to declare a variable 
	var website = "learnCodeOnline.in"  // Now here comes the lexar to save you! If we dont specify the datatype then lexar will automatically decide the datatype based upon the value of the variable .
	fmt.Println((website))
	fmt.Printf("Variable is of type: %T\n", website)

	// website = 3 // So here now  if you try to change the dataType This will show error lexar will not allow you change the datatype because it already recognized the variable as a string . 

	// no var style 
	// this is most common syntax 
	numberOfUser := 30000.0 // Here we have not used "var" keyword , not specified the dataType yet GoLang will correctly work. 
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T\n", numberOfUser)


	fmt.Printf("Variable is of type: %T\n", LoginToken) // Accessing Public variables.
    
}