package collection_test

import (
	"fmt"

	"github.com/wilhelm-murdoch/go-collection"
)

func ExampleCollection_Concat() {
	collection.New("apple", "orange", "strawberry").Concat([]string{"dog", "cat", "horse"}).Each(func(index int, item string) bool {
		fmt.Println(item)
		return false
	})

	// Output:
	// apple
	// orange
	// strawberry
	// dog
	// cat
	// horses
}

func ExampleCollection_Filter() {
	c := collection.New("apple", "orange", "strawberry").Filter(func(item string) bool {
		return item == "apple"
	})

	c.Each(func(index int, item string) bool {
		fmt.Println(item)
		return false
	})

	// Output:
	// apple
}
