package main

import "fmt"
// import "errors"

func main(){// this opening curly braces will come here only. We cant put it in the next line
	print("Monis")
	fmt.Printf("Monis %v",2) ////we can also use printf function like this
}

func print(str string){//passing a string argument and here it is type enforcer. SO i cant pass anything other than string
	fmt.Println(str)
}

// func fun(num,den)(int,int,error){ //this means two integer values will be returned and an error

	// var err error //it's initial value is nil
	// if den==0{
	// 	err=errors.New("Cant divide by zero")
	// 	return 0,0,err
	// }
	
// }