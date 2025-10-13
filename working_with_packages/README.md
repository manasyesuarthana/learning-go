# Working with Packages in Go
This section taught me in more detail about the following concepts in Go:
- Splitting code across multiple Go files
- Splitting files across multiple packages
- Importing and using custom packages

## Having more than one files with `package main`
This allows us to split functions to different files. However, there must be only **One file** with `func main()`.

- For example, message.go:
```Go
package main

import "fmt"

func printMessage() {
	message := `
	Welcome to the Go Bank!
	What do you want to do?
	1. Check Balance
	2. Deposit Money
	3. Withdraw Money
	4. Exit
	`

	fmt.Println(message)
}
```
This does not have a func main so it is not the entry point of the program. BUT we can import the function in the file with the main function:
```Go
- For example in `main.go`:
package main

import "fmt"

func main() {
    printMessage() //this is directly fetched from message.go without us needing to import it.
}
```
**But this only works if we run `go run .`**.

## Importing internal self defined packages
If we want to define a new package in our own Go workspace, we need to create to create a dedicated directory for that package. 

For example, we want to create a package dedicated for File I/O. In that case we need to create a new dir and a go file dedicated to it.
Here is an example implementation in the bank.go project:

```
.
├── balance.txt
├── bank.go
├── fileio
│   └── fileIO.go
├── go.mod
└── message.go
```
As you can see above, there is a dir named `fileio` with a `fileIO.go` file.
In the file with the main function, namely `bank.go` we can import it with the following:

```Go
import "example.com/bank-with-imports/fileio
```
We don't really need to specify the filename. We just need to specify the full path relative to our go.mod module.

**Another important note:** In order to make functions from one go file by other go files, we need to define its name starting with a capital letter. 
- Example from `bank-with-imports/fileio/fileIO.go`:
```Go
package fileio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func WriteToFile(balance float64) {
	balanceText := fmt.Sprint("Current balance: ", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func GetBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 0, errors.New("failed to find balance.txt file")
	}
	parsedBalance, err := strconv.ParseFloat(string(data)[17:], 64)
	if err != nil {
		return 0, errors.New("failed to parse balance from file")
	}
	return parsedBalance, nil

}
```

Calling it in `bank.go` will look like:
```Go
fileio.WriteToFile(...)
file.ioGetBalanceFromFile(...)
```

If a function name is not in capital, then it will not be exported/usable by other packages.

## Importing external packages
External go packages usually come with the following path `github.com/<user-name>/<package-name>`.

To import a third party go library, we need to:

1. Install the library in our local workspace:
```bash
go get github.com/<user-name>/<package-name>
```
This will download the package and add it as a **dependency** in the `go.mod` file.

2. Import it in our files:
```
import "github.com/<user-name>/<package-name>
```

3. Start using the functions from the library!

