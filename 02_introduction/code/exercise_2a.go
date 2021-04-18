package main

import (
	"fmt"
	"reflect"
)


func main() {
	// fmt.Println("Hello World")
	var x = 4
	var SomeNumber string
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(float64(x)*4.5))
	SomeNumber = "Barear"
	fmt.Println(reflect.TypeOf(SomeNumber))
	fmt.Println()
	fmt.Printf("Hi! My name is %s. I have lived in Denver for %d years. They say the weather is amazing, which is %t.\n", "Jess", 4, true)

}
