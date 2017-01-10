package main 
import "fmt"
func printInLine(a ...int){
	for _,v:=range a {
		fmt.Println(v)
	}
	beCalledFunc1(a...)
	beCalledFunc2(a)
}
func beCalledFunc1(s ...int){
	print("call in beCalledFunc1\n")
	for _,v:=range s {
		fmt.Println(v)
	}
}
func beCalledFunc2(s []int){
	print("call in beCalledFunc2\n")
	for _,v:=range s {
		fmt.Println(v)
	}
}
func main(){
	arr:=[]int{2,4,1,5}
	printInLine(arr...)
	printInLine(9,8,7,6)
}