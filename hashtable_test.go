package hashtable

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// Test basic add
// with load checks
func TestSetAndGet(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")

	assert.Equal(t, h.Get("key"), "value", "Checking set & get")
	assert.Equal(t, h.Load(), float32(0.01))
}

func TestDelete(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")
	h.Delete("key")

	assert.Equal(t, h.Get("key"), nil, "Checking set & delete")
	assert.Equal(t, h.Load(), float32(0))
}

// Test Deleting key that doesn't exist
func TestDeleteNonexistant(t *testing.T) {
	h := NewHashMap(100)
	h.Delete("key")

	assert.Equal(t, h.Delete("key"), nil, "Checking delete of nonexistant value")
	assert.Equal(t, h.Load(), float32(0))
}

// Test deleting twice
// Test Deleting key that doesn't exist
func TestDeleteTwice(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")
	h.Delete("key")

	assert.Equal(t, h.Delete("key"), nil, "Checking delete of nonexistant value")
	assert.Equal(t, h.Load(), float32(0))
}

// Test add then add again, same key
// with load checks
func TestSetAndGetAndSet(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")
	h.Set("key", "valueUpdated")

	assert.Equal(t, h.Get("key"), "valueUpdated", "Checking set & set again")
	assert.Equal(t, h.Load(), float32(0.01))
}

// Test add delete add
// with load checks
func TestSetDeleteSet(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")
	h.Delete("key")
	h.Set("key", "valueUpdated")

	assert.Equal(t, h.Get("key"), "valueUpdated", "Checking set & delete & set & get")
	assert.Equal(t, h.Load(), float32(0.01))
}

// Test add update delete
// with load checks
func TestSetUpdateDeleteSet(t *testing.T) {
	h := NewHashMap(100)
	h.Set("key", "value")
	h.Set("key", "valueUpdated")
	h.Delete("key")

	assert.Equal(t, h.Get("key"), nil, "Checking set, set, delete")
	assert.Equal(t, h.Load(), float32(0.00))
}

// Test adding entries > 1x size
// with load checks

func TestSet100(t *testing.T) {
	h := NewHashMap(100)
	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		h.Set(key, value)
	}
	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		assert.Equal(t, h.Get(key), value, "Checking batch Set & Get")
	}
	assert.Equal(t, h.Load(), float32(1.00))
}

// Test adding entries > 10x size
// with load checks

func TestSet1000(t *testing.T) {
	h := NewHashMap(100)
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		h.Set(key, value)
	}
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		assert.Equal(t, h.Get(key), value, "Checking batch Set & Get")
	}
	assert.Equal(t, h.Load(), float32(10.00))
}

// Test adding entries > 10x size and deleting
// with load checks

func TestSet1000Delete(t *testing.T) {
	h := NewHashMap(100)
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		h.Set(key, value)
	}
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		h.Delete(key)
	}
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		assert.Equal(t, h.Get(key), nil, "Checking batch Set & Delete")
	}
	assert.Equal(t, h.Load(), float32(0.00))
}

// Test constructorZero zero

func TestZeroMap(t *testing.T) {
	assert.Panics(t, func() { NewHashMap(0) }, "should panic")
}
