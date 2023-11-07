package main

import (
	"fmt"
	controlstructures "myproject/controlStructures"
	moreTypes "myproject/moretypes"
)

func main() {
	fmt.Println("Hello world!")
	controlstructures.PrintName10Times()
	fmt.Println(controlstructures.IsEven(10))
	controlstructures.SwitchNames(5)
	controlstructures.DeferStatements()
	moreTypes.Arrays()
	moreTypes.Pointers()
	moreTypes.Library()
}
