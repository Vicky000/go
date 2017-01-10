package main 
import "fmt"
func changeInfo(o *Options){
	*o.a=3
	*o.s="I am "
}
type Options struct{
	a int
	b string
}
func main(){
	var p Options
	p.a=1
	p.s="How old are you?"
	fmt.Println(p.s)
	fmt.Printf("I am %d\n",p.a)
	
	pp:=&p
	changeInfo(pp)
	print(o.s)
	print(" "+o.a)
}