package main

var a = "G"

func main(){
	n()
	m()
	n()
}

func n(){ println (a) }

func m(){
	b := "hello world"
	a = "o"
	println (a)
	println (b)
	b = "how are you"
	println (b)
}
