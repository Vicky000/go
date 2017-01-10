package main 
import "fmt"

func f(a [3]int){fmt.Println(a)}
func fp(a *[3]int){fmt.Println(a)}

func main(){
	var arr=[3]int{1,2,3}
	f(arr)
	fp(&arr)
}