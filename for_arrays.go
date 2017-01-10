package main 
import "fmt"

func main(){
	var arr1 [5]int
	
	for i:=0;i<5;i++ {
		arr1[i]=i*2
	}
	for i,a:=range arr1 {
		fmt.Printf("arr1[%d] = %d\n",i,a)
	}
	// another way to defind array
	
	a := [...]string{"a","b","c"}
	for i:=range a {
		fmt.Printf("a[%d]=%s\n",i,a[i])
	}
}