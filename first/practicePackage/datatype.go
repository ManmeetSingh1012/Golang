package practicePackage

import (
	"fmt"
	"unicode/utf8"
)

// Variables is an exported function
func Variables() {
	// declare a variable
	/*
	* var/const variable_name type = value
	* if you do not provide value, it will take default value for int = 0 , for string = "" , for bool = false (only for var not for const)
	* you cant perfom arthmatic operations on the values of  different types
	* if you divide the two int, it will give int as output : 3/2 = 1
	 */

	// int can be int64, int32, int16, int8
	// you can use simple int :- it will depend on the system arch and use int32 , and int64
	var variable_name int = 10
	fmt.Println(variable_name)

	// float : it has 2 : float32 and float64 , it does not have float
	// float32 : gives bit upto 7 decimal places
	// float64 : gives bit upto 15 decimal places

	var floating_variable float64 = 10.5
	const floating_variable2 float32 = 10.5
	fmt.Println(floating_variable, floating_variable2)

	// add 2 diff types
	var a int = 10
	var b float64 = 10.5

	// type conversion
	fmt.Print(float64(a) + b) // this will give error

	//string

	// ` ` : backticks helps to print the string as it is : hello in first and world in second line.
	var str string = `Hello 
	World`
	fmt.Println(str)

	// to print len : len will give the number of bytes
	// to print the number of characters : use utf8 package
	fmt.Println((utf8.RuneCountInString(str)))

	// rune : it is an alias for int32 , use to print unicode characters
	var r rune = 'a'
	fmt.Println(r) // unicode of a is 97

	// bool
	var boolean bool = true
	fmt.Println(boolean)

	var a1 int = 3
	var b1 int = 2
	fmt.Println(float64(a1) / float64(b1))

	// short hand
	var_name := 10
	fmt.Println(var_name)

	// multiple short hand
	x1, y, z := 10, "hello", 10.5
	fmt.Println(x1, y, z)

}
