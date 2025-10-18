package main

import (
	"fmt"
	"time"
)

// the name is capital so the struct can be used from other packages
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {

	//creating an instance of a struct
	var user1 User

	//defining contents of the data in types from functions
	user1.firstName = getUserData("=> Please enter your first name: ")
	user1.lastName = getUserData("=> Please enter your last name: ")
	user1.birthDate = getUserData("=> Please enter your birthdate (format: DD/MM/YYYY): ")
	user1.createdAt = time.Now()

	//manually or directly defining the content of a struct as we declare it
	//this is also an instance of the struct
	user2 := User{
		firstName: "some chill guy",
		lastName:  "who likes to chill",
		birthDate: "Anytime/is/chill",
		createdAt: time.Now(),
	}

	//we can also create an empty struct
	//this will automatically set all the variables in the struct to their respective null values
	user3 := User{}

	fmt.Println("*** Passing by Value: ***")
	outputUserDetails(user1)
	outputUserDetails(user2)
	outputUserDetails(user3)

	fmt.Println("*** Passing by Pointer: ***")
	outputUserDetailsWithPointer(&user1)
	outputUserDetailsWithPointer(&user2)
	outputUserDetailsWithPointer(&user3)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

// remember, since we are passing by value, we are essentially copying the value of the struct in memory
// this copy gets used in the function to be outputted, then thrown into the garbage
func outputUserDetails(passedUser User) {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", passedUser.firstName)
	fmt.Println("Last name: ", passedUser.lastName)
	fmt.Println("Date of Birth: ", passedUser.birthDate)
	fmt.Println("Created at: ", passedUser.createdAt)
	fmt.Printf("\n")
}

//now this is more efficient, no need to create another copy
//specifically for structs, Go allows us to operate structs without the need to dereference it.
//we can still dereference it, but that will take too long, notation: (*passedUser).[value]
//better to use the shortcut (without dereference notation)

func outputUserDetailsWithPointer(passedUser *User) {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", passedUser.firstName)
	fmt.Println("Last name: ", passedUser.lastName)
	fmt.Println("Date of Birth: ", passedUser.birthDate)
	fmt.Println("Created at: ", passedUser.createdAt)
	fmt.Printf("\n")
}
