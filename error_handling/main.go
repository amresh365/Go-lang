package main
import "fmt"

func divide(a,b float64)(float64,error)  {
	if b==0{
		return 0,fmt.Errorf("Detominator must not be Zero")
	}
	return a/b,nil;
}

func main()  {
	fmt.Println("This is error Handling class")
	// ans, _ := divide(10, 0) we can also write this kind
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division of two numbers:", ans)
	}
}