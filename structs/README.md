# Structs in Go
Structs are a great way to structure and group data in programming. Similar to other classic languages like C and C++, Go supports this feature and we can do a lot of awesome stuff with it.

## Defining and creating an instance of a struct
To define a struct, we need to use the `type` keyword followed by the name of the struct and finally the `struct` keyword. After that, we can define its fields. For example:
```Go
type User struct {
    firstName string
    lastName string
    birthDate string
    createdAt time.Time
}
```

To declare a struct variable, we can use multiple methods to do so:
1. Directly declaring it to the variable.
2. Use a function (global function)
3. Use a method that is a part of the struct.

### 1. Directly declaring it on the variable
```Go
newUser := User{
    firstName: "John",
    lastName: "Doe",
    birthDate: "07/08/1999"
    createdAt: time.Now()
}
```

### 2. Use a function
We can define a function that takes in the field values of the struct as arguments and return a new struct. However it is important to know that the return value of this must be a pointer of a struct, so to use the variable, we must dereference it:
```Go
func createUser(inputFirstName, inputLastName, inputBirthDate string) (*User, error) {

	if inputFirstName == "" || inputLastName == "" || inputBirthDate == "" {
		return nil, errors.New("all fields cannot be empty.")
	}
	newUser := User{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthDate: inputBirthDate,
		createdAt: time.Now(),
	}

	return &newUser, nil
}

newUser, err := createUser("John", "Doe", "07/07/1999")

//using it for operations, dereference it first:
fmt.Println(*newUser)
```

### 3. Use a struct method
This is similar to method in objects, where we can create a method for its own where the method is only accessible to that particular struct type. To achieve this, we can use **struct mutilation**, where we pass in the struct name and type before the method name:
```Go
func (user *User) outputUserDetailsMethod() {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", user.firstName)
	fmt.Println("Last name: ", user.lastName)
	fmt.Println("Date of Birth: ", user.birthDate)
	fmt.Println("Created at: ", user.createdAt)
	fmt.Printf("\n")
}

newUser := User{
    firstName: "John",
    lastName: "Doe",
    birthDate: "07/08/1999"
    createdAt: time.Now()
}

//to output user details:
newUser.outputUserDetailsMethod()
```

## Passing Structs into Functions
This is another important concept to grasp. If we want to pass a struct to a function and use the REAL value of that passed struct, we must passed a **pointer type** of that struct. Otherwise, as usual, we will just be working with a copy of it. For example, if we want to output an actual struct:
```Go
func outputUserDetailsWithPointer(passedUser *User) {
	fmt.Printf("\n === User Details === \n")
	fmt.Println("First name: ", passedUser.firstName)
	fmt.Println("Last name: ", passedUser.lastName)
	fmt.Println("Date of Birth: ", passedUser.birthDate)
	fmt.Println("Created at: ", passedUser.createdAt)
	fmt.Printf("\n")
}

newUser := User{
    firstName: "John",
    lastName: "Doe",
    birthDate: "07/08/1999"
    createdAt: time.Now()
}

outputUserDetailsWithPointer(&newUser) //this will output the actual struct
```

## Importing and Exporting structs
At some point in our projects, we would want to use structs in multiple packages. If we want our structs to be accessible by different packages, we must make it **public** and to do this, we have to make its name start with a **Capital Letter**. For example in a different package `user`, we can define this publicly accessible struct (a.k.a **exporting**):

```Go
package user

type ImportedUser struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

//With a consructor function to create an instance it:
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
```
**Additional Note**: We can also make the fields public by setting the first letter of their name to capital. But that is not good practice so we should avoid that.

Similarly, we can access/import this struct type and declare on package `main`:
```Go
package main
//... some code ...
func main() {
    // ... more code ...
    newUser, err := user.ConstructorFunction("Foreigner", "That is Foreign", "67/67/6767")
    // ... more code ...
}

```

## Embedded Structs
These are structs that have fields built on other defined structs. This is similar to **inheritance** in OOP, where we can "inherit" the features of another object in a new object type. In Go, this is possible with structs by using `Embedded structs`. For example, we have a struct type `Admin` that is built on top of `ImportedUser`:
```Go
type Admin struct {
	Email    string
	Password string
	User     ImportedUser
}
```

For a constructor, we also have to define the fields of the embedded `User` struct in the `Admin` struct:
```Go
func NewAdmin(inputEmail, inputPassword string) *Admin {
	return &Admin{
		Email:    inputEmail,
		Password: inputPassword,
		User: ImportedUser{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "Idk, who knows.",
			createdAt: time.Now(),
		},
	}
}
```

In package main, we can use it like this for example:
```Go
package main
//... some code ...
func main() {
    // ... more code ...
    admin := user.NewAdmin("admin@example.com", "admin123")
    // ... more code ...
}
```

To access the methods of the embedded struct, we need to call that struct as well:
```Go
admin.User.OutputUserDetailsMethod()
```

Output:
```
 === User Details ===
First name:  Admin
Last name:  Admin
Date of Birth:  Idk, who knows.
Created at:  2025-11-02 20:50:07.6207187 +0000 GMT m=+1.507233801
```