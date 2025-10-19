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

// methods: similars to methods in objects, we can attach a function to a struct.
// by a reciever to a function - basically the struct type it takes before the function name and after the func keyword just like:
func (user *User) outputUserDetailsMethod() {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", user.firstName)
	fmt.Println("Last name: ", user.lastName)
	fmt.Println("Date of Birth: ", user.birthDate)
	fmt.Println("Created at: ", user.createdAt)
	fmt.Printf("\n")
}

// mutating structs with methods: modifying a value inside a struct
// important: we must pass the struct as a pointer, otherwise, we are just modifying a copy of it, not the actual struct
func (user *User) makeUserAnonymous() {
	user.firstName = "Anonymous"
	user.lastName = "Anonymous"
	user.birthDate = "REDACTED"
}

// constructor methods: these are methods used to create structs - similar to constructors in objects
// for structs, make sure to return a pointer type, if not, you are just returning a copy of the struct (as usual)
func createUser(inputFirstName, inputLastName, inputBirthDate string) *User {
	newUser := User{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthDate: inputBirthDate,
		createdAt: time.Now(),
	}

	return &newUser
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

	user4 := createUser("John", "Doe", "01/01/2001") //note that this is a *User type, based on the return type of the constructor. Dereference before use.

	users := []User{user1, user2, user3, *user4}

	fmt.Println("*** (I.) Passing by Value: ***")

	for _, v := range users {
		outputUserDetails(v)
	}

	fmt.Println("*** (II.) Passing by Pointer: ***")
	// outputUserDetailsWithPointer(&user1)
	// outputUserDetailsWithPointer(&user2)
	// outputUserDetailsWithPointer(&user3)
	for _, v := range users {
		outputUserDetailsWithPointer(&v)
	}

	fmt.Println("*** (III.) Using attached Method: ***")

	for _, v := range users {
		v.outputUserDetailsMethod()
	}

	fmt.Println("*** (IV.) Make all users anonymous: ***")

	for _, v := range users {
		v.makeUserAnonymous()
		v.outputUserDetailsMethod()
	}
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
