package practicePackage

import "fmt"

func ArrayMapSliceLoops() {

	fmt.Println(("ArrayMapSliceLoops"))

	// Array
	//* array1()

	// slices
	//* slices()

	// maps
	maps()
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
	var makearray []int = make([]int, 3, 5) // make([]int, 0, 5)  0 size , cap is 5
	makearray = append(makearray, 1)
	fmt.Println(makearray) // [0 0 0 1] : because it appeand the values not replace it
}

func maps() {

	var mymap map[int]string = make(map[int]string)
	mymap[1] = "one"
	mymap[2] = "two"

	fmt.Println(mymap)
	fmt.Println(mymap)

}
