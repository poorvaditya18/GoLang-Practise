package main

import ( 
	"fmt"
	"GoLang-Practise/welcome"
	"os"
) 

func main(){
	// fmt.Println(os.Args[1:]) --> if you dont pass any argument then [] will be valid case. this wont raise error.
	fmt.Println(welcome.SayWelcome(os.Args[1:]))
}