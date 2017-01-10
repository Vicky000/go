package main 
import "fmt"

func min(a ...int) int{
	if len(a)==0 {
		return 0
	}
	min := a[0]
	for _,v:= range a {
		fmt.Printf("v=%d\n",v)
		if v < min {
			min = v
		}
	} 
	return min
}
func main(){
	x := min(1,3,2,0)
	fmt.Printf("the min is:%d\n",x)
	arr:=[]int{9,6,4,1,3}
	x = min(arr...)
	fmt.Printf("The min of arr is:%d\n",x)
}