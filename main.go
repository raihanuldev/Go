package main

import (
	"fmt"
)
const PI = 3.14
func main() {
	//loop
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	  }


	// Go Constant must be assign when it deaclre
	const fatherName string  = "John dhon\n"
	print("Father Name of constrant: ",fatherName)
	// create a varibale 
	var student1 string = "Kamrul Hasan\n"
	var a,b,c,d int = 12,13,13,14 // declare multiple varibale
	print("a: => ",a,"\n")
	print("b: => ",b,"\n")
	print("c: => ",c,"\n")
	print("d: => ",d,"\n")
	print(student1)
	fmt.Println("Hello World")
	fmt.Print("Hello Go ")
	fmt.Print("Today is Very Intersesting Day for me beacuse iam frist exploreing GoLang")
}
