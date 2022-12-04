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
	"sync/atomic"
)

type baseMap[K comparable, V any] struct {
	sync *sync.Map
	len  int32
}

func (m *baseMap[K, V]) length() int32 {
	return atomic.LoadInt32(&m.len)
}

func (m *baseMap[K, V]) exists(key K) bool {
	_, ok := m.sync.Load(key)
	return ok
}

func (m *baseMap[K, V]) setVal(key K, value V) {
	n := atomic.AddInt32(&m.len, 1)
	if n < 0 {
		panic("negative zero")
	}
	m.sync.Store(key, value)
}

func (m *baseMap[K, V]) getVal(key K) (V, bool) {
	val, ok := m.sync.Load(key)
	if val != nil {
		return val.(V), ok
	}
	var defaultValue V
	return defaultValue, false
}

func (m *baseMap[K, V]) getOrSet(key K, value V) (V, bool) {
	val, loaded := m.sync.LoadOrStore(key, value)
	if val != nil {
		return val.(V), loaded
	}
	var defaultValue V
	return defaultValue, false
}

func (m *baseMap[K, V]) getAndDelete(key K) (V, bool) {
	val, ok := m.sync.LoadAndDelete(key)
	if val != nil {
		return val.(V), ok
	}
	var defaultValue V
	return defaultValue, false
}

func (m *baseMap[K, V]) delete(key K) {
	atomic.AddInt32(&m.len, ^int32(0))
	m.sync.Delete(key)
}

func (m *baseMap[K, V]) clear() {
	m.sync.Range(func(key, value any) bool {
		if key != nil {
			atomic.AddInt32(&m.len, ^int32(0))
			m.sync.Delete(key)
			return true
		}
		return false
	})
}

func (m *baseMap[K, V]) forRange(f func(key K, val V) bool) {
	m.sync.Range(func(key, value any) bool {
		if key != nil && value != nil {
			return f(key.(K), value.(V))
		}
		return false
	})
}

func (m *baseMap[K, V]) keys() []K {
	len := atomic.LoadInt32(&m.len)
	if len <= 0 {
		return nil
	}
	keys := make([]K, len)
	i := 0
	m.sync.Range(func(key, value any) bool {
		if key != nil {
			keys[i] = key.(K)
			i++
			return true
		}
		return false
	})
	return keys
}

func (m *baseMap[K, V]) values() []V {
	len := atomic.LoadInt32(&m.len)
	if len <= 0 {
		return nil
	}
	values := make([]V, len)
	i := 0
	m.sync.Range(func(key, value any) bool {
		if value != nil {
			values[i] = value.(V)
			i++
			return true
		}
		return true
	})
	return values
}
