package collection_test

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/wilhelm-murdoch/go-collection"
)

func ExampleCollection_Sort_numeric() {
	numbers := collection.New(1, 4, 2, 3)

	fmt.Println("Unsorted:")
	numbers.Each(func(i, n int) bool {
		fmt.Print(" ", n)
		return false
	})

	fmt.Println("\nSorted:")
	numbers.Sort(func(i, j int) bool {
		left, _ := numbers.At(i)
		right, _ := numbers.At(j)
		return left < right
	}).Each(func(i, n int) bool {
		fmt.Print(" ", n)
		return false
	})

	// Output:
	// Unsorted:
	//  1 4 2 3
	// Sorted:
	//  1 2 3 4
}

func ExampleCollection_Batch() {
	type Job struct {
		Timestamp int64
		Processed bool
	}

	jobs := make([]Job, 0)
	for i := 1; i <= 100; i++ {
		jobs = append(jobs, Job{time.Now().UnixNano(), false})
	}

	c1 := collection.New(jobs...)
	c2, err := c1.Batch(func(b, j int, job Job) (Job, error) {
		job.Processed = true
		return job, nil
	}, 5)

	if err != nil {
		log.Fatal(err)
	}

	processed := c2.All(func(i int, job Job) bool {
		return job.Processed == true
	})

	fmt.Printf("processed %d/%d jobs:%v\n", c2.Length(), c1.Length(), processed)
	// Output:
	// processed 100/100 jobs:true
}

func ExampleCollection_Sort_alpha() {
	names := collection.New("wilhelm", "peter", "josh", "luke", "rob")

	fmt.Println("Unsorted:")
	fmt.Println(strings.Join(names.Items(), ","))

	names.Sort(func(i, j int) bool {
		left, _ := names.At(i)
		right, _ := names.At(j)
		return left < right
	})

	fmt.Println("Sorted:")
	fmt.Println(strings.Join(names.Items(), ","))

	// Output:
	// Unsorted:
	// wilhelm,peter,josh,luke,rob
	// Sorted:
	// josh,luke,peter,rob,wilhelm
}

func ExampleNew() {
	fruits := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot")
	fmt.Println("Fruits:", fruits.Length())

	fruits.Each(func(index int, item string) bool {
		fmt.Println("-", item)
		return false
	})

	// Output:
	// Fruits: 6
	// - apple
	// - orange
	// - strawberry
	// - cherry
	// - banana
	// - apricot
}

func ExampleCollection_Slice() {
	collection.New("apple", "orange", "strawberry").Slice(0, 2).Each(func(i int, item string) bool {
		fmt.Println(item)
		return false
	})

	// Output:
	// apple
	// orange
}

func ExampleCollection_Items() {
	c := collection.New("apple", "orange", "strawberry")

	for i, item := range c.Items() {
		fmt.Println(i, item)
	}

	// Output:
	// 0 apple
	// 1 orange
	// 2 strawberry
}

func ExampleCollection_PushDistinct() {
	c := collection.New("apple", "orange", "strawberry")

	c.PushDistinct("orange", "orange", "watermelon")

	c.Each(func(index int, item string) bool {
		fmt.Println(item)
		return false
	})

	// Output:
	// apple
	// orange
	// strawberry
	// watermelon
}

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
	// horse
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

func ExampleCollection_Contains() {
	fmt.Println(collection.New("apple", "orange", "strawberry").Contains("horse"))

	// Output:
	// false
}

func ExampleCollection_ContainsBy() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"wilhelm", 31},
		{"luke", 42},
		{"rob", 17},
		{"peter", 26},
		{"josh", 26},
	}

	fmt.Println(collection.New(people...).ContainsBy(func(i int, p Person) bool {
		return p.Age < 20
	}))

	// Output:
	// true
}

func ExampleCollection_Shift() {
	fmt.Println(collection.New("apple", "orange", "strawberry").Shift())

	// Output:
	// apple
}

func ExampleCollection_Unshift() {
	c := collection.New("apple", "orange", "strawberry")

	fmt.Println("Length Current:", c.Length())
	fmt.Println("Length New:    ", c.Unshift("horse"))

	c.Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// Length Current: 3
	// Length New:     4
	// 0 horse
	// 1 apple
	// 2 orange
	// 3 strawberry
}

func ExampleCollection_At() {
	item, ok := collection.New("apple", "orange", "strawberry").At(1)

	fmt.Println(item, ok)

	// Output:
	// orange true
}

func ExampleCollection_IsEmpty() {
	c := collection.New("lonely")

	fmt.Println(c.IsEmpty())

	c.Empty()

	fmt.Println(c.IsEmpty())

	// Output:
	// false
	// true
}

func ExampleCollection_Empty() {
	fmt.Println(collection.New("apple", "orange", "strawberry").Empty().Length())

	// Output:
	// 0
}

func ExampleCollection_Find() {
	fmt.Println(collection.New("apple", "orange", "strawberry").Find(func(i int, item string) bool {
		return item == "orange"
	}))

	// Output:
	// orange
}

func ExampleCollection_FindIndex() {
	fmt.Println(collection.New("apple", "orange", "strawberry").FindIndex(func(i int, item string) bool {
		return item == "orange"
	}))

	// Output:
	// 1
}

func ExampleCollection_Random() {
	item, ok := collection.New("apple", "orange", "strawberry").Random()

	if ok {
		fmt.Println("My random item is:", item)
	}
}

func ExampleCollection_RandomIndex() {
	index := collection.New("apple", "orange", "strawberry").RandomIndex()
	fmt.Println("My random index is:", index)
}

func ExampleCollection_LastIndexOf() {
	fmt.Println(collection.New("apple", "orange", "orange", "strawberry").LastIndexOf("orange"))

	// Output:
	// 2
}

func ExampleCollection_Reduce() {
	acc := collection.New("apple", "orange", "strawberry").Reduce(func(i int, item, accumulator string) string {
		return accumulator + item
	})

	fmt.Println(acc)

	// Output:
	// appleorangestrawberry
}

func ExampleCollection_Reverse() {
	collection.New("apple", "orange", "orange", "strawberry").Reverse().Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 strawberry
	// 1 orange
	// 2 orange
	// 3 apple
}

func ExampleCollection_Some() {
	found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
		return item == "orange"
	})

	fmt.Println("Found \"orange\"?", found)

	// Output:
	// Found "orange"? true
}

func ExampleCollection_None() {
	found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
		return item == "blackberry"
	})

	fmt.Println("Found \"blackberry\"?", found)

	// Output:
	// Found "blackberry"? false
}

func ExampleCollection_All() {
	c := collection.New("apple", "orange", "strawberry")

	fmt.Println("Contains all items?", c.All(func(i int, item string) bool {
		return c.Contains(item)
	}))

	// Output:
	// Contains all items? true
}

func ExampleCollection_Push() {
	c := collection.New("apple", "orange", "strawberry")
	fmt.Println("Collection Length:", c.Push("blueberry", "watermelon"))

	c.Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// Collection Length: 5
	// 0 apple
	// 1 orange
	// 2 strawberry
	// 3 blueberry
	// 4 watermelon
}

func ExampleCollection_Pop() {
	item, ok := collection.New("apple", "orange", "strawberry").Pop()
	fmt.Println(item, ok)

	// Output:
	// strawberry true
}

func ExampleCollection_Length() {
	fmt.Println("Collection Length:", collection.New("apple", "orange", "strawberry").Length())

	// Output:
	// Collection Length: 3
}

func ExampleCollection_Map() {
	c := collection.New("apple", "orange", "strawberry").Map(func(i int, item string) string {
		return fmt.Sprintf("The %s is yummo!", item)
	})

	c.Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 The apple is yummo!
	// 1 The orange is yummo!
	// 2 The strawberry is yummo!
}

func ExampleCollection_Each() {
	collection.New("apple", "orange", "strawberry").Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 apple
	// 1 orange
	// 2 strawberry
}

func ExampleCollection_InsertAt() {
	collection.New("apple", "orange", "strawberry").InsertAt("banana", 2).Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 apple
	// 1 orange
	// 2 banana
	// 3 strawberry
}

func ExampleCollection_InsertBefore() {
	collection.New("apple", "orange", "strawberry").InsertBefore("banana", 2).Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 apple
	// 1 banana
	// 2 orange
	// 3 strawberry
}

func ExampleCollection_InsertAfter() {
	collection.New("apple", "orange", "strawberry").InsertAfter("banana", 1).Each(func(i int, item string) bool {
		fmt.Println(i, item)
		return false
	})

	// Output:
	// 0 apple
	// 1 orange
	// 2 banana
	// 3 strawberry
}

func ExampleCollection_AtLast() {
	last, ok := collection.New("apple", "orange", "strawberry").AtLast()

	fmt.Println(last, ok)

	// Output:
	// strawberry true
}

func ExampleCollection_AtFirst() {
	first, ok := collection.New("apple", "orange", "strawberry").AtFirst()

	fmt.Println(first, ok)

	// Output:
	// apple true
}

func ExampleCollection_Count() {
	count := collection.New("apple", "orange", "orange", "strawberry").Count("orange")

	fmt.Println("Orange Count:", count)

	// Output:
	// Orange Count: 2
}

func ExampleCollection_CountBy() {
	count := collection.New("apple", "orange", "strawberry", "blueberry").CountBy(func(item string) bool {
		return strings.HasSuffix(item, "berry")
	})

	fmt.Println("Berry Types:", count)

	// Output:
	// Berry Types: 2
}

func ExampleCollection_MarshalJSON() {
	var buffer strings.Builder
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(collection.New("apple", "orange", "strawberry"))
	fmt.Println(buffer.String())

	// Output:
	// ["apple","orange","strawberry"]
}
