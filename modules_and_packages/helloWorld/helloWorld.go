package helloWorld

//different packages must belong in their respective directories
//e.g. this package is example.com/hello/helloWorld so it must be located in the helloWorld sub-dir

import "fmt"

func HelloWorld() { //if we want to export this, the name must start with an uppercase letter
	fmt.Println("Hello World!")
}