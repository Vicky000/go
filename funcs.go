package main

func main(){
	r := Abs(-3)
	print(r)
	rz := isGreater(4,3)
	print (rz)
}

func Abs(x int) int {
	if x<0 {	
		return -x
	}else{
		return x
	}
}

func isGreater(x,y int) bool {
	if x>y {
		return true
	}
	return false
}
