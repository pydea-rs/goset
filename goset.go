package goset

import "fmt"

type Set map[any]struct{}

type EmptyValue struct{}

// Add adds an item any the set and returns true if the item was added, false if it was already present.
func (set Set) Add(item any) bool {
	if _, exists := set[item]; !exists {
		set[item] = EmptyValue{}
		return true
	}
	return false
}

// Clear removes all items from the set.
func (set Set) Clear() {
	for k := range set {
		delete(set, k)
	}
}

// Has checks if an item is present in the set.
func (set Set) Has(item any) bool {
	_, exists := set[item]
	return exists
}

// Remove removes an item from the set and returns true if the item was present, false otherwise.
func (set Set) Remove(item any) bool {
	_, exists := set[item]
	if !exists {
		return false
	}
	delete(set, item)
	return true
}

// All returns a slice of all items in the set.
func (set Set) All() []any {
	items := make([]any, 0, len(set))
	for item := range set {
		items = append(items, item)
	}
	return items
}

// Count returns the number of items in the set.
func (set Set) Count() uint64 {
	return uint64(len(set))
}

// IsEmpty checks if the set is empty.
func (set Set) Empty() bool {
	return len(set) == 0
}

// Defining set string representation.
func (set Set) String() string {
	items := set.All()
	s := "{"
	for i, item := range items {
		s += fmt.Sprintf("%v", item)
		if i < len(items)-1 {
			s += ", "
		}
	}
	s += "}"
	return s
}

// NewSet creates a new empty set.
func NewSet() Set {
	return make(Set)
}
