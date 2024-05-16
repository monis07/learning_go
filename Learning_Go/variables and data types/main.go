//constant variables and basic data types
package main

import "fmt"
import "unicode/utf8"

func main(){

	var number int=32
	// var number int8
	// var number int16 = largest possible value is 32767
	// var number int32
	// var number int64 = 64 bits integer
	//we also have unsigned integers uint for storing large amount of positive integers because int8(-128,127) and uint8(0,255).

	var point float32=2.43
	//we also have float64 to store 64 bits floating point numbers
	//u have to remember that which datatype u r going to use because it will take memory accordingly
	fmt.Println(number)
	fmt.Println(point)

	//WE cant do arithmetic operations with different types of variable. If we are adding integer and float we might want to change the integer to float32 or float64 like this float32(number)+point

	//divide and modulus works same as in java or other languages

	//Strings
	/*
	strings are stored using double quotes as usual in other languages or back ticks if we have more than one line.See example below
	*/

	var mystr string ="heelloMonis"

	fmt.Println(utf8.RuneCountInString(mystr)) //find the length of the string
	fmt.Println(len(mystr)) //so we have len function which counts the no.. of bytes in the string. Go uses utf-8 encoding. Characters outside the vanilla ascii characterset are stored with more than one byte

	//boolean --   as in every other programming language
	//type inference either do var variable="Monis" or variable:="Monis"


	//Declaration
	//we can also declare variables with const but we can't change the value and we also can't just declare it, we have to initialise the value too.
}