package collection

import (
	"bytes"
	"encoding/json"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"time"
)

type Collection[T any] struct {
	items []T
}

// New returns a new collection of type T containing the specified
// items and their types. ( Chainable )
func New[T any](items ...T) *Collection[T] {
	return &Collection[T]{
		items: items,
	}
}

// Items returns the current collection's set of items.
func (c *Collection[T]) Items() []T {
	return c.items
}

// Sort sorts the collection given the provided less function. ( Chainable )
func (c *Collection[T]) Sort(less func(i, j int) bool) *Collection[T] {
	sort.Slice(c.items, less)
	return c
}

// Filter returns a new collection with items that have passed predicate check.
// ( Chainable )
func (c *Collection[T]) Filter(f func(T) bool) (out Collection[T]) {
	for _, inner := range c.items {
		if f(inner) {
			out.Push(inner)
		}
	}

	return out
}

// Batch exposes a very simple parallelized batch processor for a collection.
// Defining a `batchSize` will break the current collection into groups of
// jobs that will be processed in parallel by Goroutines managed by an error
// group. The specified function `f` will be executed for each job in each
// batch. The signature for this function is
// `func(currentBatchIndex, currentJobIndex int, job T)`. ( Chainable )
func (c *Collection[T]) Batch(f func(int, int, T), batchSize int) *Collection[T] {
	var (
		batches    [][]T
		batchCount = int(math.Ceil(float64(c.Length()) / float64(batchSize)))
		wg         sync.WaitGroup
	)

	jobsCount := c.Length()
	if batchSize > jobsCount {
		batchSize = jobsCount
	}

	offset, limit := 0, batchSize
	for i := 0; i < batchCount; i++ {
		batches = append(batches, c.items[offset:limit])
		offset, limit = limit, limit+batchSize

		if limit > c.Length() {
			limit = c.Length()
		}
	}

	for b, batch := range batches {
		wg.Add(len(batch))
		for j, t := range batch {
			go func(b int, j int, t T) {
				defer wg.Done()
				f(b, j, t)
			}(b, j, t)
		}
		wg.Wait()
	}

	return c
}

// Slice returns a new collection containing a slice of the current collection
// starting with `from` and `to` indexes. ( Chainable )
func (c *Collection[T]) Slice(from, to int) *Collection[T] {
	if from > to {
		from = to
	}

	if to > c.Length() {
		to = c.Length()
	}

	return New(c.items[from:to]...)
}

// Contains returns true if an item is present in the current collection. This
// method makes use of `reflect.DeepEqual` to ensure an absolute match. If you
// wish to check by a specific field within a slice of objects, use
// `collection.ContainsBy` instead.
func (c *Collection[T]) Contains(item T) (found bool) {
	for _, inner := range c.items {
		if reflect.DeepEqual(item, inner) {
			found = true
			break
		}
	}
	return found
}

// ContainsBy returns true if an item in the current collection matches the
// specified predicate function. This is useful if you have a slice of objects
// and you wish to check the existence of a specific field value.
func (c *Collection[T]) ContainsBy(f func(i int, item T) bool) (found bool) {
	for i, item := range c.items {
		if f(i, item) {
			found = true
			break
		}
	}
	return found
}

// PushDistinct method appends one or more distinct items to the current
// collection, returning the new length. Items that already exist within the
// current collection will be ignored. You can check for this by comparing old
// v.s. new collection lengths.
func (c *Collection[T]) PushDistinct(items ...T) int {
	for _, item := range items {
		if !c.Contains(item) {
			c.Push(item)
		}
	}

	return c.Length()
}

// Shift method removes the first item from the current collection, then
// returns that item.
func (c *Collection[T]) Shift() T {
	out := c.items[:1][0]
	c.items = c.items[1:]

	return out
}

// Unshift method appends one item to the beginning of the current collection,
// returning the new length of the collection.
func (c *Collection[T]) Unshift(item T) int {
	c.items = append([]T{item}, c.items...)
	return c.Length()
}

// At attempts to return the item associated with the specified index for the
// current collection along with a boolean value stating whether or not an item
// could be found.
func (c *Collection[T]) At(index int) (T, bool) {
	if index > (c.Length()-1) || index < 0 {
		var out T
		return out, false
	}

	return c.items[index], true
}

// IsEmpty returns a boolean value describing the empty state of the current
// collection.
func (c *Collection[T]) IsEmpty() bool {
	return c.Length() <= 0
}

// Empty will reset the current collection to zero items. ( Chainable )
func (c *Collection[T]) Empty() *Collection[T] {
	c.items = nil

	return c
}

// Find returns the first item in the provided current collectionthat satisfies
// the provided testing function. If no items satisfy the testing function,
// a <nil> value is returned.
func (c *Collection[T]) Find(f func(i int, item T) bool) (item T) {
	for i, item := range c.items {
		if found := f(i, item); found {
			return item
		}
	}

	return item
}

// FindIndex returns the index of the first item in the specified collection
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test.
func (c *Collection[T]) FindIndex(f func(i int, item T) bool) int {
	for i, item := range c.items {
		if found := f(i, item); found {
			return i
		}
	}

	return -1
}

// RandomIndex returns the index associated with a random item from the current
// collection.
func (c *Collection[T]) RandomIndex() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(c.Length() - 1)
}

// Random returns a random item from the current collection.
func (c *Collection[T]) Random() (T, bool) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return c.At(rand.Intn(c.Length()))
}

// LastIndexOf returns the last index at which a given item can be found in the
// current collection, or -1 if it is not present.
func (c *Collection[T]) LastIndexOf(item T) int {
	index := -1
	for i, inner := range c.items {
		if reflect.DeepEqual(item, inner) {
			index = i
		}
	}

	return index
}

// Reduce reduces a collection to a single value. The value is calculated by
// accumulating the result of running each item in the collection through an
// accumulator function. Each successive invocation is supplied with the return
// value returned by the previous call.
func (c *Collection[T]) Reduce(f func(i int, item, accumulator T) T) (out T) {
	for i, item := range c.items {
		out = f(i, item, out)
	}

	return out
}

// Reverse the current collection so that the first item becomes the last, the
// second item becomes the second to last, and so on. ( Chainable )
func (c *Collection[T]) Reverse() *Collection[T] {
	for i1, i2 := 0, c.Length()-1; i1 < i2; i1, i2 = i1+1, i2-1 {
		c.items[i1], c.items[i2] = c.items[i2], c.items[i1]
	}
	return c
}

// Some returns a true value if at least one item within the current collection
// resolves to true as defined by the predicate function f.
func (c *Collection[T]) Some(f func(i int, item T) bool) bool {
	for i, item := range c.items {
		if found := f(i, item); found {
			return true
		}
	}

	return false
}

// None returns a true value if no items within the current collection resolve to
// true as defined by the predicate function f.
func (c *Collection[T]) None(f func(i int, item T) bool) bool {
	count := 0
	for i, item := range c.items {
		if found := f(i, item); !found {
			count++
		}
	}

	return count == c.Length()
}

// All returns a true value if all items within the current collection resolve to
// true as defined by the predicate function f.
func (c *Collection[T]) All(f func(i int, item T) bool) bool {
	count := 0
	for i, item := range c.items {
		if found := f(i, item); found {
			count++
		}
	}

	return count == c.Length()
}

// Push method appends one or more items to the end of a collection, returning
// the new length.
func (c *Collection[T]) Push(items ...T) int {
	c.items = append(c.items, items...)
	return c.Length()
}

// Pop method removes the last item from the current collection and then
// returns that item.
func (c *Collection[T]) Pop() (out T, found bool) {
	if c.Length() == 0 {
		return
	}

	out = c.items[c.Length()-1]
	c.items = c.items[0 : c.Length()-1]

	return out, true
}

// Length returns number of items associated with the current collection.
func (c *Collection[T]) Length() int {
	return len(c.items)
}

// Map method creates to a new collection by using callback invocation result
// on each array item. On each iteration f is invoked with arguments: index and
// current item. It should return the new collection. ( Chainable )
func (c *Collection[T]) Map(f func(int, T) T) (out Collection[T]) {
	for i, item := range c.items {
		out.Push(f(i, item))
	}

	return out
}

// Each iterates through the specified list of items executes the specified
// callback on each item. This method returns the current instance of
// collection. ( Chainable )
func (c *Collection[T]) Each(f func(int, T) bool) *Collection[T] {
	for i, item := range c.items {
		if exit := f(i, item); exit {
			break
		}
	}

	return c
}

// Concat merges two slices of items. This method returns the current instance
// collection with the specified slice of items appended to it. ( Chainable )
func (c *Collection[T]) Concat(items []T) *Collection[T] {
	c.items = append(c.items, items...)
	return c
}

// InsertAt inserts the specified item at the specified index and returns the
// current collection. If the specified index is less than 0, 0 is used. If an
// index greater than the size of the collectio nis specified, c.Push is used
// instead. ( Chainable )
func (c *Collection[T]) InsertAt(item T, index int) *Collection[T] {
	if index <= 0 {
		c.Unshift(item)
		return c
	}

	if index > (c.Length() - 1) {
		c.Push(item)
		return c
	}

	c.items = append(c.items[:index+1], c.items[index:]...)
	c.items[index] = item

	return c
}

// InsertBefore inserts the specified item before the specified index and
// returns the current collection. If the specified index is less than 0,
// c.Unshift is used. If an index greater than the size of the collection is
// specified, c.Push is used instead. ( Chainable )
func (c *Collection[T]) InsertBefore(item T, index int) *Collection[T] {
	return c.InsertAt(item, index-1)
}

// InsertAfter inserts the specified item after the specified index and returns
// the current collection. If the specified index is less than 0, 0 is used. If
// an index greater than the size of the collectio nis specified, c.Push is used
// instead. ( Chainable )
func (c *Collection[T]) InsertAfter(item T, index int) *Collection[T] {
	return c.InsertAt(item, index+1)
}

// AtFirst attempts to return the first item of the collection along with a
// boolean value stating whether or not an item could be found.
func (c *Collection[T]) AtFirst() (T, bool) {
	return c.At(0)
}

// AtLast attempts to return the last item of the collection along with a
// boolean value stating whether or not an item could be found.
func (c *Collection[T]) AtLast() (T, bool) {
	return c.At(c.Length() - 1)
}

// Count counts the number of items in the collection that compare equal to value.
func (c *Collection[T]) Count(item T) (count int) {
	for _, inner := range c.items {
		if reflect.DeepEqual(inner, item) {
			count++
		}
	}

	return
}

// CountBy counts the number of items in the collection for which predicate is true.
func (c *Collection[T]) CountBy(f func(T) bool) (count int) {
	for _, item := range c.items {
		if f(item) {
			count++
		}
	}

	return count
}

// MarshalJSON implements the Marshaler interface so the current collection's
// items can be marshalled into valid JSON.
func (c *Collection[T]) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)

	if err := encoder.Encode(c.Items()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
