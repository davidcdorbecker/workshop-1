// Arrays and slices
package main

import "fmt"

func main() {
	grades := [...]int{95, 96, 90, 85, 50, 100, 96, 80}
	gradesCopy := grades
	gradesCopy[0] = 50

	fmt.Println(grades, gradesCopy)

	//a := grades[:] 	// slice of all elements
	//b := grades[3:]	// slice from 4th element to end
	//c := grades[:6]	// slice first 6 elements
	//d := grades[3:6] 	// slice the 4th, 5th and 6th element
	//
	//fmt.Println(
	//	"a: ", a, "\n",
	//	"b: ", b, "\n",
	//	"c: ", c, "\n",
	//	"d: ", d, "\n",
	//)

	//slice := make([]int, 3, 5)
	//fmt.Println(slice)
	//fmt.Println("Length: ", len(slice))
	//fmt.Println("Capacity: ", cap(slice))
}
