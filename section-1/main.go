// Values, Variables & Types
package main

import "fmt"

var (
	globalVariable = "global"
)

const globalConstant = 50

//const (
//	UP = iota
//	DOWN
//	RIGHT
//	LEFT
//)

func main() {
	var localVariable string
	var localConstant int

	localAutomaticAssign := false

	fmt.Println(
		"Global Variable: ", globalVariable, "\n",
		"Global Constant: ", globalConstant, "\n",
		"Local Variable: ", localVariable, "\n",
		"Local Constant: ", localConstant, "\n",
		"Local Automatic Assign: ", localAutomaticAssign, "\n",
	)

	//fmt.Println(
	//	UP,
	//	DOWN,
	//	RIGHT,
	//	LEFT,
	//)
}

// Default values
//0 for numeric types,
//false for the boolean type, and
//"" (the empty string) for strings.

//Types
