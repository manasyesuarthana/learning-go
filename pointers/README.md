# Pointers in Go
Pointers can be an amazing feature that we can use to achieve great feats, but it can also be something that breaks your code. It is present in classic languages like C, C++, Java, and so on.

## What is a Pointer?
It is simply a value that is used to the memory address of a variable. In this case, the variable **points** to the variable's memory address.

For example, let us create a variable and store its pointer in another variable:
```Go
package main

import "fmt"

age := 20

agePointer := &age //this is the pointer

fmt.Println(agePointer) //0xc00000a098 or some other memory.
```
In the example above, we created a variable called `age` and store its memory address in `agePointer` using the `&` ampersand symbol. In this case, `agePointer` is the pointer to `age`.

## Pointer Dereferencing
If we want to get the actual value that is pointed by a pointer, we must use pointer dereferencing with the `*` symbol:
```Go
fmt.Println("Value pointer by agePointer:", *agePointer) //20
```
This is important because if we do not dereference it, the program will return an error. In Go, we **cannot** do pointer arithmetic. 

## Passing by value vs by reference
Ah the usual material taught in classes. When passing a variable to a function, we can either pass it by **value** or **reference**.

### Passing by Value
Most programmers use this method. We directly pass the variable into the function. 

For example, let's say we have a function that converts age into adult years:
```Go
func getAdultYears(passedAge int) int {
	return passedAge - 18
}
```
Passing a variable by value simply means passing the variable directly:
```Go
age:= 20
adultAge := getAdultYears(age) //20
```
The problem is with efficiency. When we are passing a variable by value:
- Go will create a copy of that variable in memory, and this is used for the calculations in the function.
- This value is then returned.
- After it finishes, Go will throw this variable away into the garbage (garbage collection).
Because of this, a lot of uncessary work is done in the background. Instead, we can do an alternative that is more efficient. 

### Passing by Reference
In this case, we pass the pointer of the variable instead of the variable itself. This means, however, we have to adjust the function so that it can take a pointer type:
```Go
func convertToAdultYears(passedAge *int) int {
	return *passedAge - 18 //remember to dereference the pointer, since we cannot do pointer arithmetic in Go.
}
```

Passing by reference means we are passing the pointer of a variable into the function:
```Go
convertToAdultYears(&agePointer)
```
The difference lies in the fact that we do not have to create another copy of the variable. We are using this variable directly. 

However:
- The original value of the variable pointed by the pointer is NOT modified.
- If we only pass the pointer, it does not mean overwrite it.

To **actually modify** the value pointed by a pointer, we have to **mutate** it. 

## Variable Mutation
This is where we directly modify the value using pointers. In this case, the function should be written differently:
```Go
func mutateAgetoAdultYears(passedAge *int) {
	*passedAge = *passedAge - 18 //we overwrite the original value *age pointer.
}
```
In the code above, we store the arithmetic inside the pointer itself, which means we are OVERWRITING it and therefore mutating it. 

The way we call it is pretty much the same, but we can definitely see the difference in the value of the original variable after calling it:
```Go
fmt.Println("Value of age before passing its pointer to the mutating function:", *agePointer) //20
mutateAgetoAdultYears(agePointer)
fmt.Println("Value of age before after passing its pointer to the mutating function:", *agePointer) //2
```