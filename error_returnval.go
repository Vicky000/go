package main 

import (
	"errors"
	"fmt"
	"math"
)

func main(){
	fmt.Println("first example -1:")
	f,err:=MySqrt(-1)
	if err!=nil {
		fmt.Println("Error,return values:",f,err)
	}else{
		fmt.Println("return values:",f,err)
	}
	fmt.Println("second example 5:")
	if f2,err2:=MySqrt2(5);err2 != nil {
		fmt.Println("Error, return values:",f2,err2)
	}else{
		fmt.Println("return values:",f2,err2)
	}
	fmt.Println(MySqrt2(5))

}

func MySqrt( a float64) (float64,error){
	if(a<0){
		return float64(math.NaN()),errors.New("Can't deal with negative number!")
	}
	return math.Sqrt(a),nil
}

func MySqrt2(a float64) (f float64, err error){
	if (a<0){
		f=float64(math.NaN())
		err=errors.New("Can't deal with negative number!")
	}else{
		f=math.Sqrt(a)
	}
	return
}