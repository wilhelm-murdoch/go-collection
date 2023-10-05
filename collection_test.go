package collection_test

import (
	cr "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-collection"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomString(length int) string {
	b := make([]byte, length)
	cr.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func returnCollection() *collection.Collection[string] {
	return collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets", "celery", "lettuce")
}

func TestCollectionSort(t *testing.T) {
	numbers := collection.New(1, 4, 2, 3)

	numbers.Sort(func(i, j int) bool {
		left, _ := numbers.At(i)
		right, _ := numbers.At(j)
		return left < right
	})

	sorted := []int{1, 2, 3, 4}

	assert.Equal(t, numbers.Items(), sorted, "expected a sorted slice")
}

type Job struct {
	Data      string
	Processed bool
}

func TestCollectionBatch(t *testing.T) {
	jobs := make([]*Job, 0)
	for i := 0; i <= 100; i++ {
		jobs = append(jobs, &Job{randomString(15), false})
	}

	c1 := collection.New(jobs...)
	c1.Batch(func(b, j int, job *Job) {
		job.Processed = true
	}, 5)

	processed := c1.All(func(i int, job *Job) bool {
		return job.Processed == true
	})

	assert.True(t, processed, "expected all jobs to be processed")
}

func TestCollectionPushDistinct(t *testing.T) {
	c1 := returnCollection()

	l1 := c1.Length()
	l2 := c1.PushDistinct("cherry", "horse")
	assert.Equal(t, l1+1, l2, "The updated collection's length should increase by 1.")

	c2 := c1.Filter(func(s string) bool { return s == "cherry" })
	assert.Equal(t, c2.Length(), 1, "The new derivative collection could only contain a single item.")
}

func TestCollectionConcat(t *testing.T) {
	c := returnCollection()

	originalLength := c.Length()

	s := []string{"dog", "cat", "horse"}

	c.Concat(s)

	assert.Equal(t, c.Length(), (originalLength + len(s)), "The updated collection's length should equal the original length plus the length of the updated slice.")
}

func TestCollectionEach(t *testing.T) {
	c := returnCollection()

	iterations := 0
	length := c.Length()

	c.Each(func(i int, item string) bool {
		iterations++
		return false
	})

	assert.Equal(t, iterations, length, "Number of iterations must equal the lenth of the collection.")

	iterations = 0

	c.Each(func(i int, item string) bool {
		iterations++
		return true
	})

	assert.NotEqual(t, iterations, c.Length(), "Number of iterations before exiting not reached.")
}

func TestCollectionMap(t *testing.T) {
	c := returnCollection()

	append := "boop"

	cm := c.Map(func(i int, item string) string { return item + append })

	cm.Each(func(i int, item string) bool {
		assert.True(t, strings.HasSuffix(item, append), "Collection item %s should contain suffix %s.", item, append)
		return false
	})
}

func TestCollectionLength(t *testing.T) {
	c := returnCollection()

	iterations := 0
	c.Each(func(i int, item string) bool {
		iterations++
		return false
	})

	assert.Equal(t, c.Length(), iterations, "Collection length should match iteration count.")
}

func TestCollectionPop(t *testing.T) {
	c := returnCollection()

	value, ok := c.Pop()
	assert.True(t, ok, "Expected to retrieve a value, but got nothing.")
	assert.NotContains(t, value, c.Items(), "Expected value %s to not be present in collection", value)

	c.Empty()

	value, ok = c.Pop()
	assert.False(t, ok, "Expected an emptied collection, but got %s instead.", value)
}

func TestCollectionIsEmpty(t *testing.T) {
	c := returnCollection()
	assert.False(t, c.IsEmpty(), "Expected a collection with values, but got an empty one instead.")
	c.Empty()
	assert.True(t, c.IsEmpty(), "Expected an empty collection, but one items instead.")
}

func TestCollectionPush(t *testing.T) {
	c := collection.New[int]()

	s := []int{1, 2, 3}

	length := c.Push(s...)
	assert.Equal(t, length, len(s), "Expected length of %d, but got %d instead.", len(s), length)
}

func TestCollectionSome(t *testing.T) {
	c := returnCollection()

	find := "banana"
	found := c.Some(func(i int, item string) bool { return item == find })
	assert.True(t, found, "Expected value to equal %s, but got nothing instead.", find)

	find = "bananas"
	found = c.Some(func(i int, item string) bool { return item == find })
	assert.False(t, found, "Expected value equal to %s missing, but got a result instead.", find)
}

func TestCollectionReverse(t *testing.T) {
	c := returnCollection()

	old := c.Items()

	c.Reverse()

	new := c.Items()

	for i1, i2 := 0, c.Length()-1; i1 < i2; i1, i2 = i1+1, i2-1 {
		new[i1], new[i2] = new[i2], new[i1]
	}

	assert.True(t, reflect.DeepEqual(old, new), "Expected both collections to be equal, but they weren't.")
}

func TestCollectionReduce(t *testing.T) {
	t.Parallel()
	c := returnCollection()

	expected := strings.Join(c.Items(), "")

	result := c.Reduce(func(i int, item, accumulator string) string { return accumulator + item })
	assert.Equal(t, result, expected, "Expected value to equal %s, but got %s instead.", expected, result)
}

func TestCollectionLastIndexOf(t *testing.T) {
	c := collection.New("one", "two", "three", "four")

	item := "four"
	index := c.LastIndexOf(item)
	assert.Equal(t, index, 3, "Expected value %s to last appear at index %d, but got nothing.", item, index)

	item = "three"
	index = c.LastIndexOf(item)
	c.InsertAfter(item, 0)
	assert.Equal(t, c.LastIndexOf(item), 3, "Expected value %s to last appear at index %d, but got nothing.", item, index)

	index = c.LastIndexOf(item)
	c.InsertAfter(item, c.Length()-1)
	assert.Equal(t, c.LastIndexOf("four"), 4, "Expected value %s to last appear at index %d, but got nothing.", item, index)
}

func TestCollectionInsertAt(t *testing.T) {
	c := returnCollection()

	find := "squash"
	index := c.RandomIndex()

	c.InsertAt(find, index)

	found, ok := c.At(index)
	assert.True(t, ok, "Expected a value associated with random index %d, but got nothing instead.", index)
	assert.Equal(t, find, found, "Expected to find %s at index %d, but got %s instead.", find, index, found)

	c.InsertAt(find, c.Length()-1)

	found, ok = c.At(c.Length() - 1)
	assert.True(t, ok, "Expected to add value %s to the end of the collection, but but got %s instead.", find, found)
}

func TestCollectionInsertBefore(t *testing.T) {
	c := returnCollection()

	find := "squash"
	index := (c.Length() / 2)

	c.InsertBefore(find, index)

	found, ok := c.At(index - 1)
	assert.True(t, ok, "Expected a value associated with index %d, but got nothing instead.", index)
	assert.Equal(t, find, found, "Expected to find %s at index %d, but got %s instead.", find, index, found)

	index = 0

	c.InsertBefore(find, index)

	_, ok = c.At(index - 1)
	assert.False(t, ok, "Expected no values associated with index 0, but got %d instead.", found)

	first, _ := c.At(index)
	assert.Equal(t, found, first, "Expected to find %s at index %d, but got %s instead.", find, index, first)
}

func TestCollectionInsertAfter(t *testing.T) {
	c := returnCollection()

	find := "squash"
	index := c.Length() / 2

	c.InsertAfter(find, index)

	found, ok := c.At(index + 1)
	assert.True(t, ok, "Expected a value associated with index %d, but got nothing instead.", index)
	assert.Equal(t, find, found, "Expected to find %s at index %d, but got %s instead.", find, index, found)

	index = c.Length() - 1

	c.InsertAfter(find, index)

	found, ok = c.At(index + 1)
	assert.True(t, ok, "Expected no values associated with index 0, but got %d instead.", index)

	last, _ := c.At(index + 1)
	assert.Equal(t, found, last, "Expected to find %s at index %d, but got %s instead.", find, index, last)
}

func TestCollectionAtFirst(t *testing.T) {
	c := returnCollection()

	first, ok := c.AtFirst()
	assert.True(t, ok, "Expected to find a value at the beginning of the collection, but got nothing instead.")
	assert.Equal(t, first, c.Items()[0], "Expected to find value %s at the beginning of the collection, but got %s instead.", first, c.Items()[0])
}

func TestCollectionAtLast(t *testing.T) {
	c := returnCollection()

	last, ok := c.AtLast()
	assert.True(t, ok, "Expected to find a value at the end of the collection, but got nothing instead.")
	assert.Equal(t, last, c.Items()[c.Length()-1], "Expected to find value %s at the end of the collection, but got %s instead.", last, c.Items()[c.Length()-1])
}

func TestCollectionRandom(t *testing.T) {
	c := returnCollection()

	c.Each(func(i int, item string) bool {
		random, ok := c.Random()
		assert.True(t, ok, "Expected a random value, but got nothing instead.")
		assert.True(t, c.Contains(random), "Expected a random value %s to exist, but got nothing instead.", random)
		return false
	})
}

func TestCollectionRandomIndex(t *testing.T) {
	c := returnCollection()

	c.Each(func(i int, item string) bool {
		index := c.RandomIndex()
		found, ok := c.At(index)
		assert.True(t, ok, "Expected a value associated with index %d, but got nothing instead.", index)
		assert.True(t, c.Contains(found), "Expected a random value %s to exist, but got nothing instead.", found)
		return false
	})
}

func TestCollectionFindIndex(t *testing.T) {
	c := returnCollection()

	index := 3

	find, ok := c.At(index)
	assert.True(t, ok, "Expected a value associated with index %d, but got nothing instead.", index)

	found := c.FindIndex(func(i int, item string) bool { return item == find })
	assert.Equal(t, found, index, "Expected value to equal %d, but got %d instead.", index, found)

	find = "taco"
	found = c.FindIndex(func(i int, item string) bool { return item == find })
	assert.Equal(t, found, -1, "Expected value to equal %d, but got %d instead.", index, found)
}

func TestCollectionFind(t *testing.T) {
	c := returnCollection()

	find := "banana"
	found := c.Find(func(i int, item string) bool { return item == find })
	assert.NotEqual(t, len(found), 0, "Expected value to equal %s, but got nothing instead.", find)

	find = "taco"
	found = c.Find(func(i int, item string) bool { return item == find })
	assert.Equal(t, len(found), 0, "Expected value to equal nothing, but got %s instead.", find)
}

func TestCollectionAt(t *testing.T) {
	c := returnCollection()

	c.Each(func(i int, item string) bool {
		value, ok := c.At(i)
		assert.True(t, ok, "Value at index %d did not return a value.", i)
		assert.Equal(t, value, item, "Expected value to equal %d, but got %d at index %d.", value, item, i)
		return true
	})

	_, ok := c.At(1000)
	assert.False(t, ok, "Expected nil value returned, but got something else.")
}

func TestCollectionUnshift(t *testing.T) {
	c := returnCollection()

	item := "cabbage"
	c.Unshift(item)
	value, _ := c.At(0)
	assert.Equal(t, value, item, "Expected value to equal %s, but got %s instead.", item, value)
}

func TestCollectionShift(t *testing.T) {
	t.Parallel()
	c := returnCollection()

	c.Each(func(i int, item string) bool {
		l1 := c.Length()

		returned, _ := c.At(0)
		shifted := c.Shift()

		l2 := c.Length()
		assert.Equal(t, l2, (l1 - 1), "Expected collection length to %d, but got %d instead.", l2, (l1 - 1))
		assert.Equal(t, returned, shifted, "Expected value to equal %s, but got %s instead.", shifted, returned)
		return false
	})
}

func TestCollectionContains(t *testing.T) {
	c := returnCollection()

	assert.True(t, c.Contains("apple"), "Expected a value, but got nothing instead.")
	assert.False(t, c.Contains("carrots"), "Expected a value, but got nothing instead.")
}

func TestCollectionContainsBy(t *testing.T) {
	c := returnCollection()

	found := c.ContainsBy(func(i int, item string) bool {
		return strings.HasSuffix(item, "erry")
	})

	assert.True(t, found, "Expected to contain at least one value ending in `erry`, but got nothing instead.")
}

func TestCollectionSlice(t *testing.T) {
	c := returnCollection()

	s1 := c.Slice(0, 2)
	e1 := []string{"apple", "orange"}
	assert.Equal(t, s1.Length(), 2, "Expected new slice to contain 2 entries, but got %d instead.", s1.Length())
	assert.True(t, reflect.DeepEqual(e1, s1.Items()), "Expected %s, but got %s instead.", e1, s1.Items())

	s2 := c.Slice(1, 4)
	e2 := []string{"orange", "strawberry", "cherry"}
	assert.True(t, reflect.DeepEqual(e2, s2.Items()), "Expected %s, but got %s instead.", e2, s2.Items())

	s3 := c.Slice(0, 1)
	e3 := []string{"apple"}
	assert.True(t, reflect.DeepEqual(e3, s3.Items()), "Expected %s, but got %s instead.", e3, s3.Items())

	s4 := c.Slice(0, 9999)
	assert.Equal(t, c.Length(), s4.Length(), "Got a value, but expected nothing instead.")

	s5 := c.Slice(9999, 0)
	assert.False(t, s5.Length() > 0, "Got a value, but expected nothing instead.")
}

func TestCollectionFilter(t *testing.T) {
	c := returnCollection()

	out := c.Filter(func(item string) bool { return item == "strawberry" || item == "banana" })
	assert.Len(t, out.Items(), 2, "Expected a single value, but got something else.")
}

func TestCollectionCount(t *testing.T) {
	c := returnCollection()
	item := "blop"
	c.Push(item)
	c.Push(item)
	assert.Equal(t, c.Count(item), 2, "Expected %d items with value of %s, but got %d instead", 2, item, c.Count(item))
}

func TestCollectionCountBy(t *testing.T) {
	c := returnCollection()
	item := "blop"
	c.Push(item)
	c.Push(item)
	assert.Equal(t, c.CountBy(func(item string) bool { return item == "blop" }), 2, "Expected %d items with value of %s, but got %d instead", 2, item, c.Count(item))
	assert.NotEqual(t, c.CountBy(func(item string) bool { return item == "booooopsie" }), 2, "Expected %d items with value of %s, but got %d instead", 2, item, c.Count(item))
}

func TestCollectionNone(t *testing.T) {
	c := returnCollection()
	assert.True(t, c.None(func(i int, item string) bool { return !c.Contains(item) }), "Expected collection to contain no items.")
}

func TestCollectionAll(t *testing.T) {
	c := returnCollection()
	assert.True(t, c.All(func(i int, item string) bool { return c.Contains(item) }), "Expected collection to contain all items.")
}

func TestCollectionMarshalJSON(t *testing.T) {
	var buffer strings.Builder
	encoder := json.NewEncoder(&buffer)
	assert.Nil(t, encoder.Encode(returnCollection()), "Expected collection to successfully encode as valid JSON.")
	assert.Equal(t, buffer.String(), "[\"apple\",\"orange\",\"strawberry\",\"cherry\",\"banana\",\"apricot\",\"avacado\",\"beans\",\"beets\",\"celery\",\"lettuce\"]\n", "Expected collection to be marshaled into the target JSON string.")

	type Frog struct {
		Name func(name string) string
	}

	type Dog struct {
		Name string
	}

	busted := make(map[string]any, 0)

	busted["dog"] = Dog{"Spot"}
	busted["toes"] = "fingers for your feet"
	busted["frog"] = Frog{func(name string) string { return "Ribbit" }}

	buffer.Reset()
	encoder = json.NewEncoder(&buffer)
	err := encoder.Encode(collection.New(busted))
	assert.NotNil(t, err, "Expected collection marshaling to exit with an error due to unsupported mixed types.")
}
