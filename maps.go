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

type Map[K comparable, V any] interface {
	// Len returns length of map.
	Len() int32
	// Exists returns true if a key exists in the map.
	Exists(key K) bool
	// Set sets value for a key.
	Set(key K, value V)
	// SetIf sets value if the predicate function f is true.
	SetIf(key K, value V, f func(m Map[K, V]) bool) bool
	// Get returns value underlined by a key from the map.
	Get(key K) V
	// GetOk returns value and bool flag if a key in the map was found.
	GetOk(key K) (V, bool)
	// Delete deletes a key from the map.
	Delete(key K)
	// DeleteIf deletes a key from the map if the predicate function f is true.
	DeleteIf(key K, f func(m Map[K, V]) bool) bool
	// Clear clears an all keys from the map.
	Clear()
	// Keys returns all existed keys as slice in the map.
	Keys() []K
	// Values returns all existed values as slice in the map.
	Values() []V
	// Filter returns filtered pairs of keys and values from the map, if
	// predicate f is false filterting will be stopped.
	Filter(f func(key K, value V) bool) ([]K, []V)
	// Iterate iterates over each map key and value, if predicate f is false
	// iterations will be stopped.
	Iterate(f func(key K, value V) bool)
}
