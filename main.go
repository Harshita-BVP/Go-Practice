/*
	package <name>

package is a collection of common source code files

Executable package: package main is an executable package which would produce an executable file (for main.go)
on being compiled and hence, it should contain a function named main which is the entry point of the project

Reusable package: other (non-executable) helper packages can have any name other than main
*/
package main

/*
	Single package import:
		import "<package name>"

	Multiple packages import:
		import (
			"<package name>",
			"<package name>",
			"<package name>"
			...
		)
import provides access to all functionalities of a package to the current working package

fmt: is a standard go library which is used to format printing data
Go Standard Libraries are listed here: https://pkg.go.dev/std

other reusable packages can be imported in same way
*/
import "fmt"

func main() {
	fmt.Println("Hello there!")
}
