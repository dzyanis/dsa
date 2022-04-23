package main

import "fmt"

func Quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	var (
		pivot   = list[0]
		smaller = make([]int, 0, len(list))
		larger  = make([]int, 0, len(list))
	)

	for i := 1; i < len(list); i++ {
		if list[i] > pivot {
			larger = append(larger, list[i])
		} else if list[i] < pivot {
			smaller = append(smaller, list[i])
		}
	}
	return append(append(Quicksort(smaller), pivot), Quicksort(larger)...)
}

func main() {
	source := []int{9, 1, 2, 7, 5, 2, 0, -11}
	fmt.Printf("quicksort:\nwas:\t%v\ngot:\t%v\n", source, Quicksort(source))
}
