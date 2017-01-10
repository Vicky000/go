package main 

import "fmt"

var x int=10
var x2, x3 int

func main(){
	x2,x3=getX2andX3(x)
	PrintValues()
	x4:=getX2andX3_2(x)
	fmt.Print(x4)
}

func PrintValues(){
	fmt.Printf("x=%d, x2=%d, x3=%d\n",x,x2,x3)
}

func getX2andX3(input int) (int,int){
	return 2*input, 3*input

}

func getX2andX3_2(input int) (x4 int){
	//X2=2*input
	//X3=3*input
	x4=4*input
	return 

}