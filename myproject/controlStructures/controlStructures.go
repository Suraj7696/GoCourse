package controlstructures

import "fmt"

func PrintName10Times() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

func IsEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}

}

func SwitchNames(x int) {
	switch x {
	case 1:
		fmt.Println("Red")
		break
	case 2:
		fmt.Println("Blue")
		break
	case 3:
		fmt.Println("Green")
		break
	default:
		fmt.Println("White")
	}
}

func DeferStatements() {
	defer fmt.Println("This is a defered statement")
	defer fmt.Println("this is not a deferred statement")
}
