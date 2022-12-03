// Copyright (c) 2022 0x9ef
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package maps

import (
	"sync"
)

// DefaultMap is a thread-safe map that have additional helpers
// functions that give more flexibility to a user. Under the hood this
// map use sync.Map for rw-locks.
type DefaultMap[K comparable, V any] struct {
	base *baseMap[K, V]
}

// NewDefaultMap returns initialized map without capacity.
func NewDefaultMap[K comparable, V any]() DefaultMap[K, V] {
	return DefaultMap[K, V]{&baseMap[K, V]{
		sync: new(sync.Map),
	}}
}

// Len returns length of map.
func (m DefaultMap[K, V]) Len() int32 {
	return m.base.length()
}

// Exists returns true if a key exists in the map.
func (m DefaultMap[K, V]) Exists(key K) bool {
	return m.base.exists(key)
}

// Set sets value for a key.
func (m DefaultMap[K, V]) Set(key K, value V) {
	m.base.setVal(key, value)
}

// SetIf sets value if the predicate function f is true.
func (m DefaultMap[K, V]) SetIf(key K, value V, f func(m Map[K, V]) bool) bool {
	if f(m) {
		m.base.setVal(key, value)
		return true
	}
	return false
}

// Get returns value underlined by a key from the map.
func (m DefaultMap[K, V]) Get(key K) V {
	value, _ := m.base.getVal(key)
	return value
}

// GetOk returns value and bool flag if a key in the map was founded.
func (m DefaultMap[K, V]) GetOk(key K) (V, bool) {
	value, ok := m.base.getVal(key)
	if ok {
		return value, ok
	}
	var defaultValue V
	return defaultValue, false
}

// Delete deletes a key from the map.
func (m DefaultMap[K, V]) Delete(key K) {
	m.base.delete(key)
}

// DeleteIf deletes a key from the map if the predicate function f is true.
func (m DefaultMap[K, V]) DeleteIf(key K, f func(m Map[K, V]) bool) bool {
	if f(m) {
		m.base.delete(key)
		return true
	}
	return false
}

// Clear clears all keys from the map.
func (m DefaultMap[K, V]) Clear() {
	m.base.clear()
}

// Keys returns all existed keys as slice in the map.
func (m DefaultMap[K, V]) Keys() []K {
	return m.base.keys()
}

// Values returns all existed values as slice in the map.
func (m DefaultMap[K, V]) Values() []V {
	return m.base.values()
}

// Filter returns filtered pairs of keys and values from the map, if
// predicate f is false filterting will be stopped.
func (m DefaultMap[K, V]) Filter(f func(key K, value V) bool) ([]K, []V) {
	var keys []K
	var values []V
	m.base.forRange(func(key K, val V) bool {
		if f(key, val) {
			keys = append(keys, key)
			values = append(values, val)
			return true
		}
		return false
	})
	return keys, values
}

// Iterate iterates over each map key and value with iterator f function.
func (m DefaultMap[K, V]) Iterate(f func(key K, value V) bool) {
	m.base.forRange(func(key K, val V) bool {
		return f(key, val)
	})
}
