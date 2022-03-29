# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
# Usage
Import `go-collection` with the following:
```go
import (
  "github.com/wilhelm-murdoch/go-collection"
)
```
And use it like so:
```go
fruits := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot")
fmt.Println("Fruits:", fruits.Length())

fruits.Each(func(index int, item string) bool {
  fmt.Println("-", item)
  return false
})
```
Which, using the above example, will yield the following output:
```
Fruits: 6
- apple
- orange
- strawberry
- cherry
- banana
- apricot
```
# Methods

  * [New](#New)

  * [Items](#Items)

  * [Filter](#Filter)

  * [Slice](#Slice)

  * [Contains](#Contains)

  * [PushDistinct](#PushDistinct)

  * [Shift](#Shift)

  * [Unshift](#Unshift)

  * [At](#At)

  * [IsEmpty](#IsEmpty)

  * [Empty](#Empty)

  * [Find](#Find)

  * [FindIndex](#FindIndex)

  * [RandomIndex](#RandomIndex)

  * [Random](#Random)

  * [LastIndexOf](#LastIndexOf)

  * [Reduce](#Reduce)

  * [Reverse](#Reverse)

  * [Some](#Some)

  * [None](#None)

  * [All](#All)

  * [Push](#Push)

  * [Pop](#Pop)

  * [Length](#Length)

  * [Map](#Map)

  * [Each](#Each)

  * [Concat](#Concat)

  * [InsertAt](#InsertAt)

  * [InsertBefore](#InsertBefore)

  * [InsertAfter](#InsertAfter)

  * [AtFirst](#AtFirst)

  * [AtLast](#AtLast)

  * [Count](#Count)

  * [CountBy](#CountBy)

  * [returnCollection](#returnCollection)


## New
New returns a new collection of type T containing the specified
items and their types. ( Chainable )

```go

```

## Items
Items returns the current collection's set of items.

```go
c := collection.New("apple", "orange", "strawberry")

for i, item := range c.Items() {
	fmt.Println(i, item)
}
```

## Filter
Filter returns a new collection with items that have passed predicate check.
( Chainable )

```go
c := collection.New("apple", "orange", "strawberry").Filter(func(item string) bool {
	return item == "apple"
})

c.Each(func(index int, item string) bool {
	fmt.Println(item)
	return false
})
```

## Slice
Slice returns a new collection containing a slice of the current collection
starting with `from` and `to` indexes. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").Slice(0, 2).Each(func(i int, item string) bool {
	fmt.Println(item)
	return false
})
```

## Contains
Contains returns true if an item is present in the current collection.

```go
fmt.Println(collection.New("apple", "orange", "strawberry").Contains("horse"))
```

## PushDistinct
PushDistinct method appends one or more distinct items to the current
collection, returning the new length. Items that already exist within the
current collection will be ignored. You can check for this by comparing old
v.s. new collection lengths.

```go
c := collection.New("apple", "orange", "strawberry")

c.PushDistinct("orange", "orange", "watermelon")

c.Each(func(index int, item string) bool {
	fmt.Println(item)
	return false
})
```

## Shift
Shift method removes the first item from the current collection, then
returns that item.

```go
fmt.Println(collection.New("apple", "orange", "strawberry").Shift())
```

## Unshift
Unshift method appends one item to the beginning of the current collection,
returning the new length of the collection.

```go
c := collection.New("apple", "orange", "strawberry")

fmt.Println("Length Current:", c.Length())
fmt.Println("Length New:    ", c.Unshift("horse"))

c.Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## At
At attempts to return the item associated with the specified index for the
current collection along with a boolean value stating whether or not an item
could be found.

```go
item, ok := collection.New("apple", "orange", "strawberry").At(1)

fmt.Println(item, ok)
```

## IsEmpty
IsEmpty returns a boolean value describing the empty state of the current
collection.

```go
c := collection.New("lonely")

fmt.Println(c.IsEmpty())

c.Empty()

fmt.Println(c.IsEmpty())
```

## Empty
Empty will reset the current collection to zero items. ( Chainable )

```go
fmt.Println(collection.New("apple", "orange", "strawberry").Empty().Length())
```

## Find
Find returns the first item in the provided current collectionthat satisfies
the provided testing function. If no items satisfy the testing function,
a <nil> value is returned.

```go
fmt.Println(collection.New("apple", "orange", "strawberry").Find(func(i int, item string) bool {
	return item == "orange"
}))
```

## FindIndex
FindIndex returns the index of the first item in the specified collection
that satisfies the provided testing function. Otherwise, it returns -1,
indicating that no element passed the test.

```go
fmt.Println(collection.New("apple", "orange", "strawberry").FindIndex(func(i int, item string) bool {
	return item == "orange"
}))
```

## RandomIndex
RandomIndex returns the index associated with a random item from the current
collection.

```go
index := collection.New("apple", "orange", "strawberry").RandomIndex()
fmt.Println("My random index is:", index)
```

## Random
Random returns a random item from the current collection.

```go
item, ok := collection.New("apple", "orange", "strawberry").Random()

if ok {
	fmt.Println("My random item is:", item)
}
```

## LastIndexOf
LastIndexOf returns the last index at which a given item can be found in the
current collection, or -1 if it is not present.

```go
fmt.Println(collection.New("apple", "orange", "orange", "strawberry").LastIndexOf("orange"))
```

## Reduce
Reduce reduces a collection to a single value. The value is calculated by
accumulating the result of running each item in the collection through an
accumulator function. Each successive invocation is supplied with the return
value returned by the previous call.

```go
acc := collection.New("apple", "orange", "strawberry").Reduce(func(i int, item, accumulator string) string {
	return accumulator + item
})

fmt.Println(acc)
```

## Reverse
Reverse the current collection so that the first item becomes the last, the
second item becomes the second to last, and so on. ( Chainable )

```go
collection.New("apple", "orange", "orange", "strawberry").Reverse().Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## Some
Some returns a true value if at least one item within the current collection
resolves to true as defined by the predicate function f.

```go
found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
	return item == "orange"
})

fmt.Println("Found \"orange\"?", found)
```

## None
None returns a true value if no items within the current collection resolve to
true as defined by the predicate function f.

```go
found := collection.New("apple", "orange", "strawberry").Some(func(i int, item string) bool {
	return item == "blackberry"
})

fmt.Println("Found \"blackberry\"?", found)
```

## All
All returns a true value if all items within the current collection resolve to
true as defined by the predicate function f.

```go
c := collection.New("apple", "orange", "strawberry")

fmt.Println("Contains all items?", c.All(func(i int, item string) bool {
	return c.Contains(item)
}))
```

## Push
Push method appends one or more items to the end of a collection, returning
the new length.

```go
c := collection.New("apple", "orange", "strawberry")
fmt.Println("Collection Length:", c.Push("blueberry", "watermelon"))

c.Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## Pop
Pop method removes the last item from the current collection and then
returns that item.

```go
item, ok := collection.New("apple", "orange", "strawberry").Pop()
fmt.Println(item, ok)
```

## Length
Length returns number of items associated with the current collection.

```go
fmt.Println("Collection Length:", collection.New("apple", "orange", "strawberry").Length())
```

## Map
Map method creates to a new collection by using callback invocation result
on each array item. On each iteration f is invoked with arguments: index and
current item. It should return the new collection. ( Chainable )

```go
c := collection.New("apple", "orange", "strawberry").Map(func(i int, item string) string {
	return fmt.Sprintf("The %s is yummo!", item)
})

c.Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## Each
Each iterates through the specified list of items executes the specified
callback on each item. This method returns the current instance of
collection. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## Concat
Concat merges two slices of items. This method returns the current instance
collection with the specified slice of items appended to it. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").Concat([]string{"dog", "cat", "horse"}).Each(func(index int, item string) bool {
	fmt.Println(item)
	return false
})
```

## InsertAt
InsertAt inserts the specified item at the specified index and returns the
current collection. If the specified index is less than 0, 0 is used. If an
index greater than the size of the collectio nis specified, c.Push is used
instead. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").InsertAt("banana", 2).Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## InsertBefore
InsertBefore inserts the specified item before the specified index and
returns the current collection. If the specified index is less than 0,
c.Unshift is used. If an index greater than the size of the collection is
specified, c.Push is used instead. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").InsertBefore("banana", 2).Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## InsertAfter
InsertAfter inserts the specified item after the specified index and returns
the current collection. If the specified index is less than 0, 0 is used. If
an index greater than the size of the collectio nis specified, c.Push is used
instead. ( Chainable )

```go
collection.New("apple", "orange", "strawberry").InsertAfter("banana", 1).Each(func(i int, item string) bool {
	fmt.Println(i, item)
	return false
})
```

## AtFirst
AtFirst attempts to return the first item of the collection along with a
boolean value stating whether or not an item could be found.

```go
first, ok := collection.New("apple", "orange", "strawberry").AtFirst()

fmt.Println(first, ok)
```

## AtLast
AtLast attempts to return the last item of the collection along with a
boolean value stating whether or not an item could be found.

```go
last, ok := collection.New("apple", "orange", "strawberry").AtLast()

fmt.Println(last, ok)
```

## Count
Count counts the number of items in the collection that compare equal to value.

```go
count := collection.New("apple", "orange", "orange", "strawberry").Count("orange")

fmt.Println("Orange Count:", count)
```

## CountBy
CountBy counts the number of items in the collection for which predicate is true.

```go
count := collection.New("apple", "orange", "strawberry", "blueberry").CountBy(func(item string) bool {
	return strings.HasSuffix(item, "berry")
})

fmt.Println("Berry Types:", count)
```

# Benchmark
TBA
# License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.