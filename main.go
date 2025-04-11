package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a test of the msin.go file.")
	var date string ="3324";

	// var person="akg";
	person:="amresh";
	fmt.Println("The date is:", date,person)
    
	var publicVariable="123";
    var PrivateVariable="243";
	fmt.Println("private:" + publicVariable, "public:" + PrivateVariable)

	PublicFunction();
	privateFunction();
	

}
func PublicFunction() {
	fmt.Println("public function")
}

func privateFunction() {
	fmt.Println("private function")}