package practicePackage

import (
	"fmt"
	"strings"
)

func ArrayMapSliceLoops() {

	fmt.Println(("ArrayMapSliceLoops"))

	// Array
	//* array1()

	// slices
	//* slices()

	// maps
	//*maps()

	// loops
	// *loops()

	// strings
	str()
}

func array1() {

	// empty array of size 3 with all values as 0
	// array can be of same data type only.

	// always be of fixed size.

	var myarray [3]int
	myarray[0] = 123
	println(myarray[2]) // 0

	var myarray2 [3]int = [3]int{1, 2, 3} // [...]int{1, 2, 3}
	println(myarray2[2])                  // 3
}

func slices() {

	// * slices are wrapper of array , they are of dyanmic in size

	// array of size 3
	var slicearray []int = []int{1, 2, 3}
	// len : array of lenght and the cap is the capicity of the array to store the array
	fmt.Printf("the length of slice is %v and capcity is %v \n", len(slicearray), cap(slicearray))

	var smallarray []int = []int{4, 5, 6, 7}
	// appenaing the small array : it will check if we have capcity to add  3 new elements if not it will create new array with bigger cap and add the elements
	slicearray = append(slicearray, smallarray...)

	fmt.Printf("The length of slice is %v and capcity is %v \n", len(slicearray), cap(slicearray))

	// make func
	// make func : create slice array of size 3 with capcity 5
	var makearray []int = make([]int, 3, 5) // make([]int, 0, 5)  0 size , cap is 5 :- then use appeand
	makearray = append(makearray, 1)
	fmt.Println(makearray) // [0 0 0 1] : because it appeand the values not replace it
}

func maps() {

	var mymap map[int]string = make(map[int]string)
	mymap[1] = "one"
	mymap[4] = "two"

	fmt.Println(mymap)
	fmt.Println(mymap[4])

	var mymap2 map[string]int = map[string]int{"Manmeet": 21, "Chomu": 34}
	var age, ok = mymap2["Manmeet"]

	if ok {
		fmt.Printf("The age of manmeet is %v \n", age)
	} else {
		fmt.Print("Key does not exists")
	}

	//
	var age1, ok1 = mymap2["Ma"]
	if ok1 {
		fmt.Printf("The age of manmeet is %v", age1)
	} else {
		fmt.Print("Key does not exists")
	}

}

func loops() {

	var newarray map[int]string = make(map[int]string)
	newarray[1] = "kalu"
	newarray[2] = "Allu"
	newarray[3] = "Motapa"

	for key, value := range newarray {
		fmt.Printf("The key is %v and Value is %v \n", key, value)
	}

	var myarray []int = []int{1, 2, 3, 4, 5}
	for i, v := range myarray {
		fmt.Printf("The index is %v and the Value is %v \n", i, v)
	}

	// go does not have in built while loop , but you can do it initializing the the value first.

	for i := 1; i < 10; i++ {

		fmt.Println("The Value is %v", i)

	}
}

func str() {
	// in the go strings are representend int the form of bit format , go use utf-8 bit format for representation

	// when you print them then you will get ascii value

	var mystr = []rune("resmume") // rune is the other name / alias  of int32 use to print proper ascii values
	var value = mystr[0]
	fmt.Println(value)

	// concate string

	var newstr []string = []string{"a", "b", "c"}
	// using for loop and new variable we can create string abc , here new string is created , it does not modify the old array , but it is ineffficent

	// we will use string builder to concat array that use array internally , add always and the create string out of it

	var strbuilder strings.Builder
	for i := range newstr {
		strbuilder.WriteString(newstr[i])
	}
	var str = strbuilder.String()
	fmt.Printf("the string from string builder %v", str)

}
