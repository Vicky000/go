package main

import (
	"fmt"
	"unicode/utf8"
)
func main(){
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjs¤³¤ó dk"
	fmt.Printf("number of bytes of s1 : %d\n",len(s1))
	fmt.Printf("number of bytes of s2 : %d\n", len(s2))
	fmt.Printf("number of characters of s1 : %d\n", utf8.RuneCountInString(s1))
	fmt.Printf("number of characters of s2 : %d\n", utf8.RuneCountInString(s2))
	
}
