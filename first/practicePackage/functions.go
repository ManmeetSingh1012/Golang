package practicePackage

import (
	"errors"
	"fmt"
)

func Main_function() {

	simple_print("Hello World  passed to a function ")

	// mulitple values passed with return type and inserting values in x and y
	x, y := mulitple_values(2, 3, 2.5)
	fmt.Printf("The value of x is %v and y is %v", x, y)

	// with error
	result, error := multiple_values_with_error(10, 0)
	if error != nil {

		fmt.Printf("The error is %v", error.Error())
	} else {
		fmt.Printf("The result is %v", result)
	}

	// with switch case
	result2 := switch_case(2)
	fmt.Printf("The result is %v", result2)
}

func simple_print(string_value string) {

	fmt.Print(string_value)
}

// mulitple values passed with return type
func mulitple_values(a int, b int, c float64) (int, float64) {
	return a * b, c * c
}

// with error
func multiple_values_with_error(ab int, bc int) (int, error) {

	// var error error : value will be nil
	var err error
	if ab == 0 || bc == 0 {
		err = errors.New("The value of a or b is Cant be zero")
		return 0, err
	}

	return ab / bc, err

}

func switch_case(age int) string {

	// we dont use break in go

	switch {

	case age < 18:
		return "You are a minor"
	case age > 18 && age < 60:
		return "you can vote"
	default:
		return "you are a senior citizen"
	}

}
