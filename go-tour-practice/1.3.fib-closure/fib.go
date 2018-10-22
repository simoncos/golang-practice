// https://tour.go-zh.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() []int {
	list := []int{0, 1}
	return func() []int {
		current := list[len(list)-1] + list[len(list)-2]
		list = append(list, current)
		return list
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
