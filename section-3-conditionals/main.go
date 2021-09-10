package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lowerThreashold := 10
	upperThreashold := 50
	exception := 100
	number := 5

	if number > lowerThreashold && number < upperThreashold {
		fmt.Println("number is in range")
	} else if number == exception {
		fmt.Println("number is an exception")
	} else {
		if number < lowerThreashold {
			fmt.Println("number is too low")
		} else {
			fmt.Println("number is too high")
		}
	}

	multiplier := 3

	if newNum := number * multiplier; newNum > exception {
		fmt.Println("number by multiplier is greater than exception => ", newNum)
	} else {
		fmt.Println("number by multiplier is lower or equal than exception => ", newNum)
	}
	//fmt.Println(newNum)

	const (
		UP = iota
		DOWN
		RIGHT
		LEFT
	)

	cursor := rand.Intn(4)

	switch cursor {
	case UP:
		fmt.Println("UP")
	case DOWN:
		fmt.Println("DOWN")
	case LEFT:
		fmt.Println("LEFT")
	case RIGHT:
		fmt.Println("RIGHT")
	default:
		fmt.Println("UNDEFINED")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	var i interface{} = 1.0

	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}


}
