package main

import "fmt"


func simpleFunction()  {
	fmt.Println("simplr funtion")
}

func add(a,b int) (int)  {
	return a+b;
}
func main()  {
	fmt.Println("we are learning simple function in go lang")
	simpleFunction()
	ans:=add(2,3);
	fmt.Println("Add of two number:",ans);
	
}