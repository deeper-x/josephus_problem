package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []int{10, 20, 30, 40, 50, 60, 70}
	Josephus(data, 3)
}

// Josephus problem: People are standing in a circle waiting to be executed.
// Counting begins at a specified point in the circle and proceeds around the circle in a specified direction.
// After a specified number of people are skipped, the next person is executed. The procedure is repeated with the remaining people,
// starting with the next person, going in the same direction and skipping the same number of people, until only one person remains, and is freed.
func Josephus(soldiers []int, skip int) {
	r := ring.New(len(soldiers))

	for i := 0; i < len(soldiers); i++ {
		r.Value = soldiers[i]
		r = r.Next()
	}

	fmt.Printf("Starting soldier is: %v, skipping %v\n\n", r.Value, skip)

	for r.Len() > 1 {
		r.Do(func(i interface{}) {
			fmt.Printf("%v ", i.(int))
		})
		r = r.Move(skip - 1)
		fmt.Println(">> Removing", r.Next().Value)
		r.Unlink(1)
	}

	fmt.Println("\nLast soldier:", r.Value)

	// Output:
	// deeper-x@local-PC  $ go run main.go
	// Starting soldier is: 10, skipping 3

	// 10 20 30 40 50 60 70 >> Removing 40
	// 30 50 60 70 10 20 >> Removing 70
	// 60 10 20 30 50 >> Removing 30
	// 20 50 60 10 >> Removing 10
	// 60 20 50 >> Removing 60
	// 50 20 >> Removing 20

}
