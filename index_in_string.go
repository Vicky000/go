package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "Hi, I'am Marc, Hi."
	fmt.Printf("The position of Marc is: %d \n",strings.Index(str,"Marc"))
	fmt.Printf("The position of the first instance of Hi is: %d \n",strings.Index(str,"Hi"))
	fmt.Printf("The position of the last instance of Hi is: %d \n", strings.LastIndex(str,"Hi"))
	fmt.Printf("The position of Burger is: %d \n",strings.Index(str,"Burger"))

}
