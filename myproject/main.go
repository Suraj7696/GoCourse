package main

import assignment_3 "myproject/assignment_3"

func main() {
	var numbers []int
	assignment_3.PrintSlice(numbers)

	numbers = append(numbers, 0)
	assignment_3.PrintSlice(numbers)

	numbers = append(numbers, 1)
	assignment_3.PrintSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	assignment_3.PrintSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(numbers1, numbers)
	assignment_3.PrintSlice(numbers1)

	assignment_3.GoMap()

}
