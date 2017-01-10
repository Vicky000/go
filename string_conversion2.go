package main

import (
	"fmt"
	"strconv"
	//"os"
)

func main(){
	//orig := "666"
	orig := "ABC"
	//var an int
	var newS string
	
	fmt.Printf("The size of ints is : %d\n",strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("Can't convert for %s\n", orig)
		return
		// Idiom usage 1
		// return err
		
		// Idiom usage 2
		// os.Exit(1)  //return exit code
		
	}
	fmt.Printf("The integer is: %d\n",an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n",newS)

	// Idiom usage 3
	// f, err:=os.Open(filename)
	// if err != nil {
		return err
	//}

	// Idiom usage 4
	// if err:=file.Chmod(0644); err !
}
