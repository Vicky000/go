package main 
import "fmt"

func f() (ret int){
	defer func(){
		ret++
	}()
	print ("hello")
	return 1
}

func main(){
	fmt.Println(f())
}