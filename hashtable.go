package hashtable

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// This is the overarching data structure
// Broken down hierarchically into Buckets -> Entries
// Each string input hashes to the index of a bucket
type HashMap struct {
	Size    uint32
	Items   uint32
	Buckets []*Bucket
}

// A Bucket has multiple entries as it is possible
// though extremely unlikely (if the hash table is the correct size)
// that 2 different keys will hash to the same value
// this is called a "Collision"
type Bucket struct {
	Entries []*Entry
}

func NewBucket() *Bucket {
	b := new(Bucket)
	//b.Entries = make([]*Entry, 0)
	return b
}

// An entry is composed of the key and value.
// Upon accessing the entry, the key is checked
// if it maches, one can be sure its the right entry
type Entry struct {
	Key   string
	Value interface{}
}

func NewEntry(key string, value interface{}) *Entry {
	e := new(Entry)
	e.Key = key
	e.Value = value
	return e
}

// retrieves an entry with the same key
func (b *Bucket) RetrieveEntry(key string) *Entry {
	if b.Entries == nil {
		return nil
	}
	for _, entry := range b.Entries {
		if entry.Key == key {
			return entry
		}
	}
	// if that entry is not in the bucket, return nil
	return nil
}

func (b *Bucket) DeleteEntry(key string) *Entry {
	for i, entry := range b.Entries {
		if entry.Key == key {
			b.Entries = append(b.Entries[:i], b.Entries[i+1:]...)
			return entry
		}
	}
	// if that entry is not in the bucket, return nil
	return nil
}

// Constructor for the hash map
func NewHashMap(size uint32) *HashMap {
	if size <= 0 {
		panic("Cannot create HashMap size 0")
	}
	h := new(HashMap)
	h.Size = size
	h.Items = 0
	h.Buckets = make([]*Bucket, size)
	for i, _ := range h.Buckets {
		h.Buckets[i] = NewBucket()
	}
	return h
}

// Sets a value for a key in the hash map.
// Accounts for the key already having a value.
// This will overwrite any value assigned.
func (h *HashMap) Set(key string, value interface{}) bool {
	hashedKey := hash(key)
	index := hashedKey % h.Size
	if index >= 0 && index < uint32(len(h.Buckets)) {
		bucket := h.Buckets[index]
		entry := bucket.RetrieveEntry(key)
		// if there is no entry, add one and increment size
		if entry == nil {
			bucket.Entries = append(bucket.Entries, NewEntry(key, value))
			h.Items++
			return true
		} else {
			// modify the entry, leave size as is
			entry.Value = value
			return true
		}
	} else {
		//out of range error
		fmt.Println("Error in Set(): out of []Bucket range")
		return false
	}
}

// Gets a value for a key in the hash map.
// Returns false if no associated value is available
func (h *HashMap) Get(key string) interface{} {
	hashedKey := hash(key)
	index := hashedKey % h.Size
	if index >= 0 && index < uint32(len(h.Buckets)) {
		bucket := h.Buckets[index]
		entry := bucket.RetrieveEntry(key)
		if entry == nil {
			return nil
		} else {
			return entry.Value
		}
	} else {
		//out of range error
		return nil
	}
}

// Gets a value for a key in the hash map.
// Returns false if no associated value is available
func (h *HashMap) Delete(key string) interface{} {
	hashedKey := hash(key)
	index := hashedKey % h.Size
	if index >= 0 && index < uint32(len(h.Buckets)) {
		bucket := h.Buckets[index]
		value := bucket.DeleteEntry(key)
		if value == nil {
			return nil
		} else {
			h.Items--
			return value
		}
	} else {
		//out of range error
		return nil
	}
}

// returns load factor
func (h *HashMap) Load() float32 {
	if h.Size == 0 {
		return -1 //signals an error
	}
	return float32(h.Items) / float32(h.Size)
}
