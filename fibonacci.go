package main 
import "fmt"
func main(){
	result := 0
	for i := 0; i<=10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i,result)
	}
}
func fibonacci(i int) (s int){
	if i<=1 {
		s=1
	} else {
		s=fibonacci(i-1)+fibonacci(i-2)
	}
	return
}