package main

import "fmt"

func main(){
LABEL1:
	for i:=0;i<=5;i++{
		for j:=0;j<=5;j++{
				if j==4{
				break LABEL1
				}
				fmt.Printf("i=%d, j=%d\n",i,j)
		}
		
	}
}