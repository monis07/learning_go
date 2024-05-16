//Structure of a go program
/*
so we need to understand two things about the structure of go code
1. Modules -- Collection of packages
2.Packages-- Package is just a collection of go files
*/

// 1. go is a statically typed language which means that we have to declare variable types explicitly or they have to be infered
package main

import "fmt"

//main function will be created only where the package name is main
func main(){
	var variable string="Monis" //declaring a variable of type string
	variable1:="Azeem" //here putting colon so that it can be infered what it is 

	fmt.Println(variable)
	fmt.Println(variable1)

	//2. go is Strongly typed language. We cant add two different types of variables like in JavaScript.Because of this compiler make us do through checking of our code and fix bugs

	//3. Go also comes with a compiler which translates our code to machine code(binary file) which makes it faster than other languages like python and ruby and we can execute that executable file on any machine without any dependencies

	//4. Fast compile time--  If u run a loop in go till million times it takes about 50ms in go compared to python which takes around 6s
	//5. Built in concurrency -- Done in Go with Goroutines
	//6. Simplicity and Garbage collection

}

