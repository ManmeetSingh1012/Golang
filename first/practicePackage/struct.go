package practicePackage

import (
	"fmt"
)

// Structs are custom types

func StructInterface() {
	// Call the struct function
	Struct()
}

type techStack struct {
	name        string
	job_opening int
	language
}

type language struct {
	language_name []string
}

// attached function to the e techstack
func (e techStack) calculateSalary() int {
	// Assuming job_opening represents the number of openings and each opening pays 100000 units
	return e.job_opening * 1000
}

type nontech struct {
	name        string
	job_opening int
}

func (e nontech) calculateSalary() int {
	// Assuming job_opening represents the number of openings and each opening pays 100000 units
	return e.job_opening * 1000
}

// to check richness we cant pass all stack we use interface for that :- where we have predefine values
type stacks interface {
	calculateSalary() int
}

func Rich(e stacks) bool {

	if e.calculateSalary() > 5000 {
		fmt.Printf("your are rich ")
		return true
	} else {
		fmt.Print("Poor Kid ")
		return false
	}
}

func Struct() {
	// Correctly initialize the language_name slice
	var student1 techStack = techStack{"Backend", 10, language{[]string{"c++", "java"}}}
	fmt.Printf("The Salary is %v\n", student1.language.language_name)

	var student2 techStack = techStack{"BlockChain", 2, language{[]string{"Solana", "Rust"}}}
	fmt.Printf("The Salary is %v\n", student2.calculateSalary())

	// check richness
	var techguy techStack = techStack{"Mlops", 1, language{[]string{"python"}}}
	fmt.Println(Rich(techguy))

	var nonttechguy nontech = nontech{"sales", 1000}
	fmt.Println(Rich(nonttechguy))

}
