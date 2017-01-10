package main 
import "fmt"
func main(){
	arr := [3]float64{7.0,8.5,9.1}
	sum := Sum(&arr)
	fmt.Printf("sum of arr[] is: %f\n",sum)
}
func Sum(a *[3]float64)(sum float64){
	for _,v := range a {
		sum+=v
	}
	return
}