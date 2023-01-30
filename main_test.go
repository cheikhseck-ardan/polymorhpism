package main

import "testing"

func TestAdd(t *testing.T) {

	t.Run("Test String", func(t *testing.T) {

		expecting := "Foo Baz"
		output := Add[string]("Foo", "Baz")

		if output != expecting {
			t.Fatal(
				"Test failed, expecting :",
				expecting,
				"got",
				output,
			)
		}
	})

	t.Run("Test Int", func(t *testing.T) {

		expecting := 100
		output := Add[int](50, 50)

		if output != expecting {
			t.Fatal(
				"Test failed, expecting :",
				expecting,
				"got",
				output,
			)
		}
	})
}
