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

func newDefaultMap[V any](b *testing.B) Map[int, V] {
	m := NewDefaultMap[int, V]()
	for i := 0; i < b.N; i++ {
		var defaultValue V
		m.Set(i, defaultValue)
	}
	return m
}

func BenchmarkDefaultMapLenInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Len()
	}
}

func BenchmarkDefaultMapExistsInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Exists(i)
	}
}

func BenchmarkDefaultMapSetInt(b *testing.B) {
	m := newDefaultMap[int](b)
	for i := 0; i < b.N; i++ {
		m.Set(i, 0)
	}
}

func BenchmarkDefaultMapSetIfInt(b *testing.B) {
	m := newDefaultMap[int](b)
	for i := 0; i < b.N; i++ {
		m.SetIf(i, 0, func(m Map[int, int]) bool {
			return m.Get(b.N) == 0
		})
	}
}

func BenchmarkDefaultMapGetInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(i)
	}
}

func BenchmarkDefaultMapDeleteInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

func BenchmarkDefaultMapDeleteIfInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.DeleteIf(i, func(m Map[int, int]) bool {
			return i == b.N/2
		})
	}
}

func BenchmarkDefaultMapKeysInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Keys()
	}
}

func BenchmarkDefaultMapValuesInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Values()
	}
}

func BenchmarkDefaultMapFilterInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Filter(func(key, value int) bool {
			return key == b.N/2 && value == 0
		})
	}
}

func BenchmarkDefaultMapIterateInt(b *testing.B) {
	m := newDefaultMap[int](b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Iterate(func(key, value int) bool {
			return true
		})
	}
}
