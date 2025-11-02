package user

import (
	"errors"
	"fmt"
	"time"
)

// creating a struct that can be accessed by other packages: using a capital letter for the first letter of the name.
// this does not only apply for the struct name, but for the attributes as well
// however, that should not be possible as it can pose a "risk", instead we should use a constructor function to define the values
type ImportedUser struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (user *ImportedUser) OutputUserDetailsMethod() {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", user.firstName)
	fmt.Println("Last name: ", user.lastName)
	fmt.Println("Date of Birth: ", user.birthDate)
	fmt.Println("Created at: ", user.createdAt)
	fmt.Printf("\n")
}

// mutating structs with methods: modifying a value inside a struct
// important: we must pass the struct as a pointer, otherwise, we are just modifying a copy of it, not the actual struct
func (user *ImportedUser) MakeUserAnonymous() {
	user.firstName = "Anonymous"
	user.lastName = "Anonymous"
	user.birthDate = "REDACTED"
}

// here is a constructor function for the ImportedUser class:
func ConstructorFunction(inputFirstName, inputLastName, inputBirthDate string) (*ImportedUser, error) {
	if inputFirstName == "" || inputLastName == "" || inputBirthDate == "" {
		return nil, errors.New("all fields must not be empty.")
	}

	return &ImportedUser{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthDate: inputBirthDate,
		createdAt: time.Now(),
	}, nil
}
