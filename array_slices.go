package main 
import "fmt"
func main(){
	var arr1 [6]int
	var slice1 []int = arr1[1:5]
	var slice2 []int
	for i:=0;i<len(arr1);i++ {
		arr1[i]=i 
	}
	
	for i:=0;i<len(slice1);i++ {
		fmt.Printf("slice1[%d] = %d\n",i,slice1[i])
	}
	
	fmt.Printf("the length of arr1 is %d\n",len(arr1))
	fmt.Printf("the length of slice1 is %d\n",len(slice1))
	fmt.Printf("the capacity of slice1 is %d\n",cap(slice1))
	
	slice1=slice1[:cap(slice1)]
	//slice1=slice1[-1:]
	for i:=0;i<len(slice1);i++ {
		fmt.Printf("slice1[%d] = %d\n",i,slice1[i])
	}
	
	fmt.Printf("the length of slice1 is %d\n",len(slice1))
	fmt.Printf("the capacity of slice1 is %d\n",cap(slice1))
	
	slice2=slice1[1:]
	for i:=0;i<len(slice2);i++ {
		fmt.Printf("slice2[%d] = %d\n",i,slice2[i])
	}
	
	b:=[]byte{'g','o','l','a','n','g'}
	s1:=b[1:4]
	s2:=b[:2]
	s3:=b[2:]
	s4:=b[:]
	fmt.Print("\ns1:\n")
	for _,v:=range s1 {
		fmt.Printf("%c ",v)
	}
	fmt.Print("\ns2:\n")
	for _,v:=range s2 {
		fmt.Printf("%c ",v)
	}
	fmt.Print("\ns3:\n")
	for _,v:=range s3 {
		fmt.Printf("%c ",v)
	}
	fmt.Print("\ns4:\n")
	for _,v:=range s4 {
		fmt.Printf("%c ",v)
	}
}