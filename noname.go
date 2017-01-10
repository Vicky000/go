package main

func main(){
	fv:=func(){ print ("hello world")}
	fv()
	switch t:=fv.(type){
	case string:
		println("is string:%s",t)
	default:
		println(t)
	}
}

