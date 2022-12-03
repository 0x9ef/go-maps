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

// UniqueMap is a thread-safe map that operates of zero-value
// keys and have an additional helper functions that give more flexibility
// to a user.
type UniqueMap[K comparable] struct {
	baseMap *baseMap[K, struct{}]
}

var uniqueMapZeroValue = struct{}{}

// NewUniqueMap returns initialized with make map.
func NewUniqueMap[K comparable]() UniqueMap[K] {
	return UniqueMap[K]{
		&baseMap[K, struct{}]{
			sync: new(sync.Map),
		},
	}
}

// Len returns length of map.
func (m UniqueMap[K]) Len() int32 {
	return m.baseMap.length()
}

// Exists returns true if a key exists in the map.
func (m UniqueMap[K]) Exists(key K) bool {
	return m.baseMap.exists(key)
}

// Set sets value for a key.
func (m UniqueMap[K]) Set(key K) {
	m.baseMap.setVal(key, uniqueMapZeroValue)
}

// SetIf sets value if the predicate function f is true.
func (m UniqueMap[K]) SetIf(key K, f func(m UniqueMap[K]) bool) bool {
	if f(m) {
		m.baseMap.setVal(key, uniqueMapZeroValue)
		return true
	}
	return false
}

// Delete deletes a key from the map.
func (m UniqueMap[K]) Delete(key K) {
	m.baseMap.delete(key)
}

// Clear clears all keys from the map.
func (m UniqueMap[K]) Clear() {
	m.baseMap.clear()
}

// Keys returns all existed keys as slice in the map.
func (m UniqueMap[K]) Keys() []K {
	return m.baseMap.keys()
}
