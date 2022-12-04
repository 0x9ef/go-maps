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
	"testing"
)

func newUniqueMap[V any](n int) UniqueMap[int] {
	m := NewUniqueMap[int]()
	for i := 0; i < n; i++ {
		m.Set(i)
	}
	return m
}

func BenchmarkUniqueMapLenInt(b *testing.B) {
	m := newUniqueMap[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Len()
	}
}

func BenchmarkUniqueMapExistsInt(b *testing.B) {
	m := newUniqueMap[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Exists(i)
	}
}

func BenchmarkUniqueMapSetInt(b *testing.B) {
	m := NewUniqueMap[int]()
	for i := 0; i < b.N; i++ {
		m.Set(i)
	}
}

func BenchmarkMapSetIf(b *testing.B) {
	m := newUniqueMap[int](b.N)
	for i := 0; i < b.N; i++ {
		m.SetIf(i, func(m UniqueMap[int]) bool {
			return m.Exists(b.N / 2)
		})
	}
}

func BenchmarkUniqueMapDeleteInt(b *testing.B) {
	m := newUniqueMap[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

func BenchmarkUniqueMapClearInt(b *testing.B) {
	m := newUniqueMap[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Clear()
	}
}

func BenchmarkUniqueMapKeysInt(b *testing.B) {
	m := newUniqueMap[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Keys()
	}
}
