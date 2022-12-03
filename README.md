# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

A generic collection for Go with a few convenient methods. 

There are far more comprehensive modules out there, but this one works quite well for my purposes.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
# Reference

* [New](#Function-New)
* [Items](#Function-Items)
* [Sort](#Function-Sort)
* [Filter](#Function-Filter)
* [Batch](#Function-Batch)
* [Slice](#Function-Slice)
* [Contains](#Function-Contains)
* [ContainsBy](#Function-ContainsBy)
* [PushDistinct](#Function-PushDistinct)
* [Shift](#Function-Shift)
* [Unshift](#Function-Unshift)
* [At](#Function-At)
* [IsEmpty](#Function-IsEmpty)
* [Empty](#Function-Empty)
* [Find](#Function-Find)
* [FindIndex](#Function-FindIndex)
* [RandomIndex](#Function-RandomIndex)
* [Random](#Function-Random)
* [LastIndexOf](#Function-LastIndexOf)
* [Reduce](#Function-Reduce)
* [Reverse](#Function-Reverse)
* [Some](#Function-Some)
* [None](#Function-None)
* [All](#Function-All)
* [Push](#Function-Push)
* [Pop](#Function-Pop)
* [Length](#Function-Length)
* [Map](#Function-Map)
* [Each](#Function-Each)
* [Concat](#Function-Concat)
* [InsertAt](#Function-InsertAt)
* [InsertBefore](#Function-InsertBefore)
* [InsertAfter](#Function-InsertAfter)
* [AtFirst](#Function-AtFirst)
* [AtLast](#Function-AtLast)
* [Count](#Function-Count)
* [CountBy](#Function-CountBy)
* [MarshalJSON](#Function-MarshalJSON)


### Function `New`
* `func New[T any](items ...T) *Collection[T]` [#](collection.go#L20)
* `collection.go:20:24` [#](collection.go#L20-L24)

New returns a new collection of type T containing the specified items and their types. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fruits := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot")
    fmt.Println("Fruits:", fruits.Length())
    
    fruits.Each(func(index int, item string) bool {
    	fmt.Println("-", item)
    	return false
    })
}
```
```go
// Output:
// Fruits: 6
// - apple
// - orange
// - strawberry
// - cherry
// - banana
// - apricot
```
### Function `Items`
* `func (c *Collection[T]) Items() []T` [#](collection.go#L27)
* `collection.go:27:29` [#](collection.go#L27-L29)

Items returns the current collection's set of items. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry")
    
    for i, item := range c.Items() {
    	fmt.Println(i, item)
    }
}
```
```go
// Output:
// 0 apple
// 1 orange
// 2 strawberry
```
### Function `Sort`
* `func (c *Collection[T]) Sort(less func(i, j int) bool) *Collection[T]` [#](collection.go#L32)
* `collection.go:32:35` [#](collection.go#L32-L35)

Sort sorts the collection given the provided less function. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
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
}
```
```go
// Output:
// Unsorted:
//  1 4 2 3
// Sorted:
//  1 2 3 4
```
```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
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
}
```
```go
// Output:
// Unsorted:
// wilhelm,peter,josh,luke,rob
// Sorted:
// josh,luke,peter,rob,wilhelm
```
### Function `Filter`
* `func (c *Collection[T]) Filter(f func(T) bool) (out Collection[T])` [#](collection.go#L39)
* `collection.go:39:47` [#](collection.go#L39-L47)

Filter returns a new collection with items that have passed predicate check. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry").Filter(func(item string) bool {
    	return item == "apple"
    })
    
    c.Each(func(index int, item string) bool {
    	fmt.Println(item)
    	return false
    })
}
```
```go
// Output:
// apple
```
### Function `Batch`
* `func (c *Collection[T]) Batch(f func(int, int, T), batchSize int) *Collection[T]` [#](collection.go#L55)
* `collection.go:55:89` [#](collection.go#L55-L89)

Batch exposes a very simple parallelized batch processor for a collection. Defining a `batchSize` will break the current collection into groups of jobs that will be processed in parallel by Goroutines managed by an error group. The specified function `f` will be executed for each job in each batch. The signature for this function is `func(currentBatchIndex, currentJobIndex int, job T)`. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    type Job struct {
    	Timestamp int64
    	Processed bool
    }
    
    jobs := make([]*Job, 0)
    for i := 1; i <= 100; i++ {
    	jobs = append(jobs, &Job{time.Now().UnixNano(), false})
    }
    
    c1 := collection.New(jobs...)
    c1.Batch(func(b, j int, job *Job) {
    	job.Processed = true
    }, 5)
    
    processed := c1.All(func(i int, job *Job) bool {
    	return job.Processed == true
    })
    
    fmt.Printf("processed %d/%d jobs:%v\n", c1.Length(), c1.Length(), processed)
}
```
```go
// Output:
// processed 100/100 jobs:true
```
### Function `Slice`
* `func (c *Collection[T]) Slice(from, to int) *Collection[T]` [#](collection.go#L93)
* `collection.go:93:103` [#](collection.go#L93-L103)

Slice returns a new collection containing a slice of the current collection starting with `from` and `to` indexes. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").Slice(0, 2).Each(func(i int, item string) bool {
    	fmt.Println(item)
    	return false
    })
}
```
```go
// Output:
// apple
// orange
```
### Function `Contains`
* `func (c *Collection[T]) Contains(item T) (found bool)` [#](collection.go#L109)
* `collection.go:109:117` [#](collection.go#L109-L117)

Contains returns true if an item is present in the current collection. This method makes use of `reflect.DeepEqual` to ensure an absolute match. If you wish to check by a specific field within a slice of objects, use `collection.ContainsBy` instead. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "strawberry").Contains("horse"))
}
```
```go
// Output:
// false
```
### Function `ContainsBy`
* `func (c *Collection[T]) ContainsBy(f func(i int, item T) bool) (found bool)` [#](collection.go#L122)
* `collection.go:122:130` [#](collection.go#L122-L130)

ContainsBy returns true if an item in the current collection matches the specified predicate function. This is useful if you have a slice of objects and you wish to check the existence of a specific field value. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
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
}
```
```go
// Output:
// true
```
### Function `PushDistinct`
* `func (c *Collection[T]) PushDistinct(items ...T) int` [#](collection.go#L136)
* `collection.go:136:144` [#](collection.go#L136-L144)

PushDistinct method appends one or more distinct items to the current collection, returning the new length. Items that already exist within the current collection will be ignored. You can check for this by comparing old v.s. new collection lengths. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry")
    
    c.PushDistinct("orange", "orange", "watermelon")
    
    c.Each(func(index int, item string) bool {
    	fmt.Println(item)
    	return false
    })
}
```
```go
// Output:
// apple
// orange
// strawberry
// watermelon
```
### Function `Shift`
* `func (c *Collection[T]) Shift() T` [#](collection.go#L148)
* `collection.go:148:153` [#](collection.go#L148-L153)

Shift method removes the first item from the current collection, then returns that item. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "strawberry").Shift())
}
```
```go
// Output:
// apple
```
### Function `Unshift`
* `func (c *Collection[T]) Unshift(item T) int` [#](collection.go#L157)
* `collection.go:157:160` [#](collection.go#L157-L160)

Unshift method appends one item to the beginning of the current collection, returning the new length of the collection. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry")
    
    fmt.Println("Length Current:", c.Length())
    fmt.Println("Length New:    ", c.Unshift("horse"))
    
    c.Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// Length Current: 3
// Length New:     4
// 0 horse
// 1 apple
// 2 orange
// 3 strawberry
```
### Function `At`
* `func (c *Collection[T]) At(index int) (T, bool)` [#](collection.go#L165)
* `collection.go:165:172` [#](collection.go#L165-L172)

At attempts to return the item associated with the specified index for the current collection along with a boolean value stating whether or not an item could be found. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    item, ok := collection.New("apple", "orange", "strawberry").At(1)
    
    fmt.Println(item, ok)
}
```
```go
// Output:
// orange true
```
### Function `IsEmpty`
* `func (c *Collection[T]) IsEmpty() bool` [#](collection.go#L176)
* `collection.go:176:178` [#](collection.go#L176-L178)

IsEmpty returns a boolean value describing the empty state of the current collection. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("lonely")
    
    fmt.Println(c.IsEmpty())
    
    c.Empty()
    
    fmt.Println(c.IsEmpty())
}
```
```go
// Output:
// false
// true
```
### Function `Empty`
* `func (c *Collection[T]) Empty() *Collection[T]` [#](collection.go#L181)
* `collection.go:181:185` [#](collection.go#L181-L185)

Empty will reset the current collection to zero items. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "strawberry").Empty().Length())
}
```
```go
// Output:
// 0
```
### Function `Find`
* `func (c *Collection[T]) Find(f func(i int, item T) bool) (item T)` [#](collection.go#L190)
* `collection.go:190:198` [#](collection.go#L190-L198)

Find returns the first item in the provided current collectionthat satisfies the provided testing function. If no items satisfy the testing function, a <nil> value is returned. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "strawberry").Find(func(i int, item string) bool {
    	return item == "orange"
    }))
}
```
```go
// Output:
// orange
```
### Function `FindIndex`
* `func (c *Collection[T]) FindIndex(f func(i int, item T) bool) int` [#](collection.go#L203)
* `collection.go:203:211` [#](collection.go#L203-L211)

FindIndex returns the index of the first item in the specified collection that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "strawberry").FindIndex(func(i int, item string) bool {
    	return item == "orange"
    }))
}
```
```go
// Output:
// 1
```
### Function `RandomIndex`
* `func (c *Collection[T]) RandomIndex() int` [#](collection.go#L215)
* `collection.go:215:218` [#](collection.go#L215-L218)

RandomIndex returns the index associated with a random item from the current collection. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    index := collection.New("apple", "orange", "strawberry").RandomIndex()
    fmt.Println("My random index is:", index)
}
```

### Function `Random`
* `func (c *Collection[T]) Random() (T, bool)` [#](collection.go#L221)
* `collection.go:221:224` [#](collection.go#L221-L224)

Random returns a random item from the current collection. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    item, ok := collection.New("apple", "orange", "strawberry").Random()
    
    if ok {
    	fmt.Println("My random item is:", item)
    }
}
```

### Function `LastIndexOf`
* `func (c *Collection[T]) LastIndexOf(item T) int` [#](collection.go#L228)
* `collection.go:228:237` [#](collection.go#L228-L237)

LastIndexOf returns the last index at which a given item can be found in the current collection, or -1 if it is not present. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println(collection.New("apple", "orange", "orange", "strawberry").LastIndexOf("orange"))
}
```
```go
// Output:
// 2
```
### Function `Reduce`
* `func (c *Collection[T]) Reduce(f func(i int, item, accumulator T) T) (out T)` [#](collection.go#L243)
* `collection.go:243:249` [#](collection.go#L243-L249)

Reduce reduces a collection to a single value. The value is calculated by accumulating the result of running each item in the collection through an accumulator function. Each successive invocation is supplied with the return value returned by the previous call. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    acc := collection.New("apple", "orange", "strawberry").Reduce(func(i int, item, accumulator string) string {
    	return accumulator + item
    })
    
    fmt.Println(acc)
}
```
```go
// Output:
// appleorangestrawberry
```
### Function `Reverse`
* `func (c *Collection[T]) Reverse() *Collection[T]` [#](collection.go#L253)
* `collection.go:253:258` [#](collection.go#L253-L258)

Reverse the current collection so that the first item becomes the last, the second item becomes the second to last, and so on. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "orange", "strawberry").Reverse().Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 strawberry
// 1 orange
// 2 orange
// 3 apple
```
### Function `Some`
* `func (c *Collection[T]) Some(f func(i int, item T) bool) bool` [#](collection.go#L262)
* `collection.go:262:270` [#](collection.go#L262-L270)

Some returns a true value if at least one item within the current collection resolves to true as defined by the predicate function f. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
    	return item == "orange"
    })
    
    fmt.Println("Found \"orange\"?", found)
}
```
```go
// Output:
// Found "orange"? true
```
### Function `None`
* `func (c *Collection[T]) None(f func(i int, item T) bool) bool` [#](collection.go#L274)
* `collection.go:274:283` [#](collection.go#L274-L283)

None returns a true value if no items within the current collection resolve to true as defined by the predicate function f. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
    	return item == "blackberry"
    })
    
    fmt.Println("Found \"blackberry\"?", found)
}
```
```go
// Output:
// Found "blackberry"? false
```
### Function `All`
* `func (c *Collection[T]) All(f func(i int, item T) bool) bool` [#](collection.go#L287)
* `collection.go:287:296` [#](collection.go#L287-L296)

All returns a true value if all items within the current collection resolve to true as defined by the predicate function f. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry")
    
    fmt.Println("Contains all items?", c.All(func(i int, item string) bool {
    	return c.Contains(item)
    }))
}
```
```go
// Output:
// Contains all items? true
```
### Function `Push`
* `func (c *Collection[T]) Push(items ...T) int` [#](collection.go#L300)
* `collection.go:300:303` [#](collection.go#L300-L303)

Push method appends one or more items to the end of a collection, returning the new length. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry")
    fmt.Println("Collection Length:", c.Push("blueberry", "watermelon"))
    
    c.Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// Collection Length: 5
// 0 apple
// 1 orange
// 2 strawberry
// 3 blueberry
// 4 watermelon
```
### Function `Pop`
* `func (c *Collection[T]) Pop() (out T, found bool)` [#](collection.go#L307)
* `collection.go:307:316` [#](collection.go#L307-L316)

Pop method removes the last item from the current collection and then returns that item. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    item, ok := collection.New("apple", "orange", "strawberry").Pop()
    fmt.Println(item, ok)
}
```
```go
// Output:
// strawberry true
```
### Function `Length`
* `func (c *Collection[T]) Length() int` [#](collection.go#L319)
* `collection.go:319:321` [#](collection.go#L319-L321)

Length returns number of items associated with the current collection. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    fmt.Println("Collection Length:", collection.New("apple", "orange", "strawberry").Length())
}
```
```go
// Output:
// Collection Length: 3
```
### Function `Map`
* `func (c *Collection[T]) Map(f func(int, T) T) (out Collection[T])` [#](collection.go#L326)
* `collection.go:326:332` [#](collection.go#L326-L332)

Map method creates to a new collection by using callback invocation result on each array item. On each iteration f is invoked with arguments: index and current item. It should return the new collection. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    c := collection.New("apple", "orange", "strawberry").Map(func(i int, item string) string {
    	return fmt.Sprintf("The %s is yummo!", item)
    })
    
    c.Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 The apple is yummo!
// 1 The orange is yummo!
// 2 The strawberry is yummo!
```
### Function `Each`
* `func (c *Collection[T]) Each(f func(int, T) bool) *Collection[T]` [#](collection.go#L337)
* `collection.go:337:345` [#](collection.go#L337-L345)

Each iterates through the specified list of items executes the specified callback on each item. This method returns the current instance of collection. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 apple
// 1 orange
// 2 strawberry
```
### Function `Concat`
* `func (c *Collection[T]) Concat(items []T) *Collection[T]` [#](collection.go#L349)
* `collection.go:349:352` [#](collection.go#L349-L352)

Concat merges two slices of items. This method returns the current instance collection with the specified slice of items appended to it. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").Concat([]string{"dog", "cat", "horse"}).Each(func(index int, item string) bool {
    	fmt.Println(item)
    	return false
    })
}
```
```go
// Output:
// apple
// orange
// strawberry
// dog
// cat
// horse
```
### Function `InsertAt`
* `func (c *Collection[T]) InsertAt(item T, index int) *Collection[T]` [#](collection.go#L358)
* `collection.go:358:373` [#](collection.go#L358-L373)

InsertAt inserts the specified item at the specified index and returns the current collection. If the specified index is less than 0, 0 is used. If an index greater than the size of the collectio nis specified, c.Push is used instead. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").InsertAt("banana", 2).Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 apple
// 1 orange
// 2 banana
// 3 strawberry
```
### Function `InsertBefore`
* `func (c *Collection[T]) InsertBefore(item T, index int) *Collection[T]` [#](collection.go#L379)
* `collection.go:379:381` [#](collection.go#L379-L381)

InsertBefore inserts the specified item before the specified index and returns the current collection. If the specified index is less than 0, c.Unshift is used. If an index greater than the size of the collection is specified, c.Push is used instead. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").InsertBefore("banana", 2).Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 apple
// 1 banana
// 2 orange
// 3 strawberry
```
### Function `InsertAfter`
* `func (c *Collection[T]) InsertAfter(item T, index int) *Collection[T]` [#](collection.go#L387)
* `collection.go:387:389` [#](collection.go#L387-L389)

InsertAfter inserts the specified item after the specified index and returns the current collection. If the specified index is less than 0, 0 is used. If an index greater than the size of the collectio nis specified, c.Push is used instead. ( Chainable ) 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    collection.New("apple", "orange", "strawberry").InsertAfter("banana", 1).Each(func(i int, item string) bool {
    	fmt.Println(i, item)
    	return false
    })
}
```
```go
// Output:
// 0 apple
// 1 orange
// 2 banana
// 3 strawberry
```
### Function `AtFirst`
* `func (c *Collection[T]) AtFirst() (T, bool)` [#](collection.go#L393)
* `collection.go:393:395` [#](collection.go#L393-L395)

AtFirst attempts to return the first item of the collection along with a boolean value stating whether or not an item could be found. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    first, ok := collection.New("apple", "orange", "strawberry").AtFirst()
    
    fmt.Println(first, ok)
}
```
```go
// Output:
// apple true
```
### Function `AtLast`
* `func (c *Collection[T]) AtLast() (T, bool)` [#](collection.go#L399)
* `collection.go:399:401` [#](collection.go#L399-L401)

AtLast attempts to return the last item of the collection along with a boolean value stating whether or not an item could be found. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    last, ok := collection.New("apple", "orange", "strawberry").AtLast()
    
    fmt.Println(last, ok)
}
```
```go
// Output:
// strawberry true
```
### Function `Count`
* `func (c *Collection[T]) Count(item T) (count int)` [#](collection.go#L404)
* `collection.go:404:412` [#](collection.go#L404-L412)

Count counts the number of items in the collection that compare equal to value. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    count := collection.New("apple", "orange", "orange", "strawberry").Count("orange")
    
    fmt.Println("Orange Count:", count)
}
```
```go
// Output:
// Orange Count: 2
```
### Function `CountBy`
* `func (c *Collection[T]) CountBy(f func(T) bool) (count int)` [#](collection.go#L415)
* `collection.go:415:423` [#](collection.go#L415-L423)

CountBy counts the number of items in the collection for which predicate is true. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    count := collection.New("apple", "orange", "strawberry", "blueberry").CountBy(func(item string) bool {
    	return strings.HasSuffix(item, "berry")
    })
    
    fmt.Println("Berry Types:", count)
}
```
```go
// Output:
// Berry Types: 2
```
### Function `MarshalJSON`
* `func (c *Collection[T]) MarshalJSON() ([]byte, error)` [#](collection.go#L427)
* `collection.go:427:436` [#](collection.go#L427-L436)

MarshalJSON implements the Marshaler interface so the current collection's items can be marshalled into valid JSON. 

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
    var buffer strings.Builder
    encoder := json.NewEncoder(&buffer)
    encoder.Encode(collection.New("apple", "orange", "strawberry"))
    fmt.Println(buffer.String())
}
```
```go
// Output:
// ["apple","orange","strawberry"]
```


Documentation generated by [Gadget](https://github.com/wilhelm-murdoch/go-gadget).

# License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
