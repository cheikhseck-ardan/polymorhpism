package main

import (
	"fmt"
	"log"
)

func main() {

	log.Println(Add[string]("Foo", "Bar"))
	log.Println(Add[int](40, 50))
}

func Add[T any](a any, b any) T {

	var result T

	switch a := a.(type) {
	case string:
		txt := fmt.Sprintf(
			"%s %s",
			a,
			b,
		)
		result = any(txt).(T)
	case int:
		result = any(a + b.(int)).(T)
	}

	return result
}
