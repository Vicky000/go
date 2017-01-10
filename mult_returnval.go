package main
import "fmt"
func Caculator(a,b int) (sum int, multiplication int, d int){
	sum=a+b
	multiplication=a*b
	d=a-b
	return
}

func Caculator_2(a,b int) (int, int, int){
	return a+b,a*b,a-b
}

func main(){
sum,m,d:=Caculator(2,1)
sum2,m2,d2:=Caculator_2(2,1)
fmt.Println ("Sum=",sum,"Product=",m,"Diff=",d)
fmt.Println ("Sum=",sum2,"Product=",m2,"Diff=",d2)
}