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
  "github.com/wilhelm-murdoch/lo"
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
* [Filter](#Filter) 
* [Slice](#Slice) 
* [Contains](#Contains) 
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

  
## New
New returns a new collection of type T containing the specified
items and their types.
 
## Filter
Filter returns a new collection with items that have passed predicate check.
 
## Slice
Slice returns a new slice of the current collection starting with `from` and
`to` indexes.
 
## Contains
Contains returns true if an item is present in the current collection.
 
## Shift
Shift method removes the first item from the current collection, then
returns that item.
 
## Unshift
Unshift method appends one item to the beginning of the current collection,
returning the new length of the collection.
 
## At
At attempts to return the item associated with the specified index for the
current collection along with a boolean value stating whether or not an item
could be found.
 
## IsEmpty
IsEmpty returns a boolean value describing the empty state of the current
collection.
 
## Empty
Empty will reset the current collection to zero items.
 
## Find
Find returns the first item in the provided current collectionthat satisfies
the provided testing function. If no items satisfy the testing function,
a &lt;nil&gt; value is returned.
 
## FindIndex
FindIndex returns the index of the first item in the specified collection
that satisfies the provided testing function. Otherwise, it returns -1,
indicating that no element passed the test.
 
## RandomIndex
RandomIndex returns the index associated with a random item from the current
collection.
 
## Random
Random returns a random item from the current collection.
 
## LastIndexOf
LastIndexOf returns the last index at which a given item can be found in the
current collection, or -1 if it is not present.
 
## Reduce
Reduce reduces a collection to a single value. The value is calculated by
accumulating the result of running each item in the collection through an
accumulator function. Each successive invocation is supplied with the return
value returned by the previous call.
 
## Reverse
Reverse the current collection so that the first item becomes the last, the
second item becomes the second to last, and so on.
 
## Some
Some returns a true value if at least one item within the current collection
resolves to true as defined by the predicate function f.
 
## None
None returns a true value if no items within the current collection resolve to
true as defined by the predicate function f.
 
## All
All returns a true value if all items within the current collection resolve to
true as defined by the predicate function f.
 
## Push
Push method appends one or more items to the end of a collection, returning
the new length.
 
## Pop
Pop method removes the last item from the current collection and then
returns that item.
 
## Length
Length returns number of items associated with the current collection.
 
## Map
Map method creates to a new collection by using callback invocation result
on each array item. On each iteration f is invoked with arguments: index and
current item. It should return the new collection.
 
## Each
Each iterates through the specified list of items executes the specified
callback on each item. This method returns the current instance of collection.
 
## Concat
Concat merges two slices of items. This method returns the current instance
collection with the specified slice of items appended to it.
 
## InsertAt
InsertAt inserts the specified item at the specified index and returns the
current collection. If the specified index is less than 0, 0 is used. If an
index greater than the size of the collectio nis specified, c.Push is used
instead.
 
## InsertBefore
InsertBefore inserts the specified item before the specified index and returns
the current collection. If the specified index is less than 0, 0 is used. If
an index greater than the size of the collectio nis specified, c.Push is used
instead.
 
## InsertAfter
InsertAfter inserts the specified item after the specified index and returns
the current collection. If the specified index is less than 0, 0 is used. If
an index greater than the size of the collectio nis specified, c.Push is used
instead.
 
## AtFirst
AtFirst attempts to return the first item of the collection along with a
boolean value stating whether or not an item could be found.
 
## AtLast
AtLast attempts to return the last item of the collection along with a
boolean value stating whether or not an item could be found.
 
## Count
Count counts the number of items in the collection that compare equal to value.
 
## CountBy
CountBy counts the number of items in the collection for which predicate is true.
   
