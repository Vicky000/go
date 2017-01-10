package main

var a = "G"

func main(){
	n()
	m()
	n()
}

func n(){ print(a) }

func m(){
	a := "o"
	print (a)
}
