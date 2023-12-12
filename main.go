package main

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

func TypeOf[T any](v T) string {
	return fmt.Sprintf("%T", v)
}

type List []int

func Search[S ~[]E, E constraints.Integer](s S, t E) int {
	for idx, s := range s {
		if t == s {
			return idx
		}
	}

	return -1
}

type MyString string

// Example for defining constaints
type PrintableType interface {
	int | string
}

func Print[T PrintableType](x T) {
	fmt.Println(x)
}

func main() {

	// Example 1: type safety is ensured at compile time as
	// opposed to reflection during runtime
	x := 42
	// Using reflection to determine the type of 'x'
	t := reflect.TypeOf(x)

	fmt.Println("Type of x with reflection:", t)
	fmt.Println("Type of x with generics:", TypeOf(x))

	// Example 2: Type inference
	d := Search([]int{1, 2, 3}, 3)        // no type given,compiler can deduce it
	r := Search[[]int]([]int{1, 2, 3}, 3) // type given
	fmt.Printf("WithType %d and withoutType %d", r, d)

	// Example 3: Constraints and underlying type
	// var customStr MyString
	// customStr = "g"
	// Print(customStr) // will throw error at compile time
	// How to fix it?

	// Example 4: Constraint type inference
	list := List{1, 2, 3, 4}
	idx := Search(list, 4)
	fmt.Println(idx)

}
