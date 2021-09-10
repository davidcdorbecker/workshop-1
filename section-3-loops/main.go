package main

import "fmt"

func main() {
	for i:= 0 ; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 1; i < 5; i, j = i+1, j+2{
		fmt.Println(i, j)
	}

	//while loop?
	i:=0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	////infinite loop
	//for {
	//	fmt.Println("aiuda")
	//}

	slice := []int{1, 2, 3}

	for i, value := range slice {
		fmt.Println(i, value)
	}
}
