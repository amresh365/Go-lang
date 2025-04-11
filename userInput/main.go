package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
fmt.Println("What is your Name:")
var name string;
//  fmt.Scan(&name);
//  fmt.Println("Your Name is"+name)

//  This is use to take input with white space 
 reader:=bufio.NewReader(os.Stdin)
 name,_=reader.ReadString('\n')
 fmt.Println("Hello my name is "+name)

}