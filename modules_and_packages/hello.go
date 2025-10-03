package main //this is a special package in go, it is the entry point of the program

//the main package would always have to be located in the main directory

import (
	"example.com/hello/helloWorld" //we are importing a package called helloWorld, located in the helloWorld subdir
)

func main(){
	helloWorld.HelloWorld()
}