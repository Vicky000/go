package main
import "fmt"

func main() {
// di yi ti
	for i:=0;i<15;i++ {
		fmt.Printf("The counter is %d\n", i)
	}
// di er ti
	i:= 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 { goto START }
}
