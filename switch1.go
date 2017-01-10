package main

import "fmt"

func main(){
	var num1 int = 99
	switch num1 {
		case 98,99:
			fmt.Println("is 98")
			fallthrough
		case 100:
			fmt.Println("is 100")
		default:
			fmt.Println(" is between 98 and 100")	
	}
}
