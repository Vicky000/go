package main 
import "fmt"
func main(){
	f1()
}

func f1(){
	
	print("at the top of f1\n")
	defer f2()
	defer a()
	
	print("at the bottom of f1\n")
	
}
func f2(){
	print ("call f2\n")
}

func a(){
	i:=0
	defer fmt.Println(i)
	i++
	return}