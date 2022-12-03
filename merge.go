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

type Merger[K comparable, V any] interface {
	// Merge merges all maps to the one combined map.
	Merge(collection ...Map[K, V]) Map[K, V]
	// MergeUnique merges only unique elements to the map.
	MergeUnique(collection ...Map[K, V]) Map[K, V]
}

// Merge merges all maps to the one combined map.
func Merge[K comparable, V any](collection ...Map[K, V]) Map[K, V] {
	mergeMap := NewDefaultMap[K, V]()
	for _, next := range collection {
		next.Filter(func(key K, value V) bool {
			mergeMap.base.setVal(key, value)
			return true
		})
	}
	return mergeMap
}

// MergeUnique merges only unique elements to the map.
func MergeUnique[K comparable, V any](collection ...Map[K, V]) Map[K, V] {
	mergeMap := NewDefaultMap[K, V]()
	for _, next := range collection {
		next.Filter(func(key K, val V) bool {
			unique := mergeMap.base.exists(key)
			if !unique {
				mergeMap.base.setVal(key, val)
				return true
			}
			return false
		})
	}
	return mergeMap
}
