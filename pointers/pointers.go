package main

import "fmt"

func main() {
	age := 20          // regular variable, not a pointer
	agePointer := &age //this is a pointer, stores the address of age

	fmt.Println("Given age in adult years:", getAdultYears(age)) //we are passing by the variable
	fmt.Println("Address of Age:", agePointer)

	//passing the value by pointer or reference
	newAge := convertToAdultYears(agePointer) //this does not modify the original age value, however.
	fmt.Println("Given age in adult years (passing with pointer):", newAge)

	//value mutation

	//pointer dereferencing: add an asterisk again --> to get the value that is pointed by a pointer
	fmt.Println("Value of age before passing its pointer to the mutating function:", *agePointer)
	mutateAgetoAdultYears(agePointer)
	fmt.Println("Value of age before after passing its pointer to the mutating function:", *agePointer)

}

//when we pass in the value by variable, Go will create a copy of this value
//this copy will continue to get used until the function finishes execution, after that, it is garbage collected.
func getAdultYears(passedAge int) int {
	return passedAge - 18
}

//this function is different from the above, it allows us to pass a pointer.
//in this case, no copy of the variable is created, we are passing the actual variable and modifying it through the function
//however, this does not modify the ACTUAL value of the pointer, yet, in that case, we have to MUTATE it.
func convertToAdultYears(passedAge *int) int {
	return *passedAge - 18 //remember to dereference the pointer, since we cannot do pointer arithmetic in Go.
}

//now we are also passing the value by reference, but this time, we are modifying the actual variable.
//in other words, we are mutating it.
func mutateAgetoAdultYears(passedAge *int) {
	*passedAge = *passedAge - 18 //we overwrite the original value *age pointer.
}
