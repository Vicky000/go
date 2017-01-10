package main 

import "strings"
import "fmt"

func main(){
	for i:=1;i<26;i++ {
		for j:=0;j<i;j++ {
			print("G")
		}
		print("\n")
	}

// 2
	print (" 2. use strings \n")
	s:= "G"
	for i:=1;i<26;i++ {
		S:=strings.Repeat(s,i)
		fmt.Println(S)
	}
// 3
	print ("3. use string concatenation")
	for i:=1;i<26;i++ {
		println(s)
		s += "G"
	}
	
}
