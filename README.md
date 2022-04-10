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
* [Filter](#Function-Filter)
* [Slice](#Function-Slice)
* [Contains](#Function-Contains)
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
* `func New[T any](items ...T) *Collection[T]` [#](collection.go#L17)
* `collection.go:17:21` [#](collection.go#L17-L21)

New returns a new collection of type T containing the specifieditems and their types. ( Chainable )

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
* `func (c *Collection[T]) Items() []T` [#](collection.go#L24)
* `collection.go:24:26` [#](collection.go#L24-L26)

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
### Function `Filter`
* `func (c *Collection[T]) Filter(f func(T) bool) (out Collection[T])` [#](collection.go#L30)
* `collection.go:30:38` [#](collection.go#L30-L38)

Filter returns a new collection with items that have passed predicate check.( Chainable )

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
### Function `Slice`
* `func (c *Collection[T]) Slice(from, to int) *Collection[T]` [#](collection.go#L42)
* `collection.go:42:52` [#](collection.go#L42-L52)

Slice returns a new collection containing a slice of the current collectionstarting with `from` and `to` indexes. ( Chainable )

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
* `func (c *Collection[T]) Contains(item T) (found bool)` [#](collection.go#L55)
* `collection.go:55:63` [#](collection.go#L55-L63)

Contains returns true if an item is present in the current collection.

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
### Function `PushDistinct`
* `func (c *Collection[T]) PushDistinct(items ...T) int` [#](collection.go#L69)
* `collection.go:69:77` [#](collection.go#L69-L77)

PushDistinct method appends one or more distinct items to the currentcollection, returning the new length. Items that already exist within thecurrent collection will be ignored. You can check for this by comparing oldv.s. new collection lengths.

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
* `func (c *Collection[T]) Shift() T` [#](collection.go#L81)
* `collection.go:81:86` [#](collection.go#L81-L86)

Shift method removes the first item from the current collection, thenreturns that item.

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
* `func (c *Collection[T]) Unshift(item T) int` [#](collection.go#L90)
* `collection.go:90:93` [#](collection.go#L90-L93)

Unshift method appends one item to the beginning of the current collection,returning the new length of the collection.

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
* `func (c *Collection[T]) At(index int) (T, bool)` [#](collection.go#L98)
* `collection.go:98:105` [#](collection.go#L98-L105)

At attempts to return the item associated with the specified index for thecurrent collection along with a boolean value stating whether or not an itemcould be found.

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
* `func (c *Collection[T]) IsEmpty() bool` [#](collection.go#L109)
* `collection.go:109:111` [#](collection.go#L109-L111)

IsEmpty returns a boolean value describing the empty state of the currentcollection.

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
* `func (c *Collection[T]) Empty() *Collection[T]` [#](collection.go#L114)
* `collection.go:114:118` [#](collection.go#L114-L118)

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
* `func (c *Collection[T]) Find(f func(i int, item T) bool) (item T)` [#](collection.go#L123)
* `collection.go:123:131` [#](collection.go#L123-L131)

Find returns the first item in the provided current collectionthat satisfiesthe provided testing function. If no items satisfy the testing function,a <nil> value is returned.

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
* `func (c *Collection[T]) FindIndex(f func(i int, item T) bool) int` [#](collection.go#L136)
* `collection.go:136:144` [#](collection.go#L136-L144)

FindIndex returns the index of the first item in the specified collectionthat satisfies the provided testing function. Otherwise, it returns -1,indicating that no element passed the test.

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
* `func (c *Collection[T]) RandomIndex() int` [#](collection.go#L148)
* `collection.go:148:151` [#](collection.go#L148-L151)

RandomIndex returns the index associated with a random item from the currentcollection.

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
* `func (c *Collection[T]) Random() (T, bool)` [#](collection.go#L154)
* `collection.go:154:157` [#](collection.go#L154-L157)

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
* `func (c *Collection[T]) LastIndexOf(item T) int` [#](collection.go#L161)
* `collection.go:161:170` [#](collection.go#L161-L170)

LastIndexOf returns the last index at which a given item can be found in thecurrent collection, or -1 if it is not present.

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
* `func (c *Collection[T]) Reduce(f func(i int, item, accumulator T) T) (out T)` [#](collection.go#L176)
* `collection.go:176:182` [#](collection.go#L176-L182)

Reduce reduces a collection to a single value. The value is calculated byaccumulating the result of running each item in the collection through anaccumulator function. Each successive invocation is supplied with the returnvalue returned by the previous call.

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
* `func (c *Collection[T]) Reverse() *Collection[T]` [#](collection.go#L186)
* `collection.go:186:191` [#](collection.go#L186-L191)

Reverse the current collection so that the first item becomes the last, thesecond item becomes the second to last, and so on. ( Chainable )

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
* `func (c *Collection[T]) Some(f func(i int, item T) bool) bool` [#](collection.go#L195)
* `collection.go:195:203` [#](collection.go#L195-L203)

Some returns a true value if at least one item within the current collectionresolves to true as defined by the predicate function f.

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
* `func (c *Collection[T]) None(f func(i int, item T) bool) bool` [#](collection.go#L207)
* `collection.go:207:216` [#](collection.go#L207-L216)

None returns a true value if no items within the current collection resolve totrue as defined by the predicate function f.

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
* `func (c *Collection[T]) All(f func(i int, item T) bool) bool` [#](collection.go#L220)
* `collection.go:220:229` [#](collection.go#L220-L229)

All returns a true value if all items within the current collection resolve totrue as defined by the predicate function f.

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
* `func (c *Collection[T]) Push(items ...T) int` [#](collection.go#L233)
* `collection.go:233:236` [#](collection.go#L233-L236)

Push method appends one or more items to the end of a collection, returningthe new length.

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
* `func (c *Collection[T]) Pop() (out T, found bool)` [#](collection.go#L240)
* `collection.go:240:249` [#](collection.go#L240-L249)

Pop method removes the last item from the current collection and thenreturns that item.

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
* `func (c *Collection[T]) Length() int` [#](collection.go#L252)
* `collection.go:252:254` [#](collection.go#L252-L254)

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
* `func (c *Collection[T]) Map(f func(int, T) T) (out Collection[T])` [#](collection.go#L259)
* `collection.go:259:265` [#](collection.go#L259-L265)

Map method creates to a new collection by using callback invocation resulton each array item. On each iteration f is invoked with arguments: index andcurrent item. It should return the new collection. ( Chainable )

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
* `func (c *Collection[T]) Each(f func(int, T) bool) *Collection[T]` [#](collection.go#L270)
* `collection.go:270:278` [#](collection.go#L270-L278)

Each iterates through the specified list of items executes the specifiedcallback on each item. This method returns the current instance ofcollection. ( Chainable )

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
* `func (c *Collection[T]) Concat(items []T) *Collection[T]` [#](collection.go#L282)
* `collection.go:282:285` [#](collection.go#L282-L285)

Concat merges two slices of items. This method returns the current instancecollection with the specified slice of items appended to it. ( Chainable )

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
* `func (c *Collection[T]) InsertAt(item T, index int) *Collection[T]` [#](collection.go#L291)
* `collection.go:291:306` [#](collection.go#L291-L306)

InsertAt inserts the specified item at the specified index and returns thecurrent collection. If the specified index is less than 0, 0 is used. If anindex greater than the size of the collectio nis specified, c.Push is usedinstead. ( Chainable )

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
* `func (c *Collection[T]) InsertBefore(item T, index int) *Collection[T]` [#](collection.go#L312)
* `collection.go:312:314` [#](collection.go#L312-L314)

InsertBefore inserts the specified item before the specified index andreturns the current collection. If the specified index is less than 0,c.Unshift is used. If an index greater than the size of the collection isspecified, c.Push is used instead. ( Chainable )

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
* `func (c *Collection[T]) InsertAfter(item T, index int) *Collection[T]` [#](collection.go#L320)
* `collection.go:320:322` [#](collection.go#L320-L322)

InsertAfter inserts the specified item after the specified index and returnsthe current collection. If the specified index is less than 0, 0 is used. Ifan index greater than the size of the collectio nis specified, c.Push is usedinstead. ( Chainable )

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
* `func (c *Collection[T]) AtFirst() (T, bool)` [#](collection.go#L326)
* `collection.go:326:328` [#](collection.go#L326-L328)

AtFirst attempts to return the first item of the collection along with aboolean value stating whether or not an item could be found.

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
* `func (c *Collection[T]) AtLast() (T, bool)` [#](collection.go#L332)
* `collection.go:332:334` [#](collection.go#L332-L334)

AtLast attempts to return the last item of the collection along with aboolean value stating whether or not an item could be found.

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
* `func (c *Collection[T]) Count(item T) (count int)` [#](collection.go#L337)
* `collection.go:337:344` [#](collection.go#L337-L344)

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
* `func (c *Collection[T]) CountBy(f func(T) bool) (count int)` [#](collection.go#L347)
* `collection.go:347:354` [#](collection.go#L347-L354)

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
* `func (c *Collection[T]) MarshalJSON() ([]byte, error)` [#](collection.go#L358)
* `collection.go:358:367` [#](collection.go#L358-L367)

MarshalJSON implements the Marshaler interface so the current collection'sitems can be marshalled into valid JSON.

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

# License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
