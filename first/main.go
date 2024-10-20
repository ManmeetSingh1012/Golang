// package main : belongs to main package
package main

// package fmt
// they way to import packages
import (
	// "first/testpackage"
	"first/practicePackage"
	"fmt"
)

func main() {
	fmt.Println("A le chak main aa gaya : ")

	// // to check the function from test.go
	// testpackage.Test()

	// run this to get idea about variables
	// * practicePackage.Variables()

	// run this to get idea about functions
	// * practicePackage.Main_function()

	// run this to get idea about array , map , slice , loops
	practicePackage.ArrayMapSliceLoops()
}
