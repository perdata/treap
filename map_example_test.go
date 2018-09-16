// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package treap_test

import (
	"fmt"
	"github.com/perdata/treap"
	"math/rand"
)

type pair struct {
	key, value interface{}
}

type comparer func(kv1, kv2 pair) int

func (c comparer) Compare(v1, v2 interface{}) int {
	return c(v1.(pair), v2.(pair))
}

type Map struct {
	*treap.Node
	treap.Comparer
}

func NewMap(keyCompare treap.Comparer) Map {
	c := comparer(func(v1, v2 pair) int {
		return keyCompare.Compare(v1.key, v2.key)
	})
	return Map{nil, c}
}

func (m Map) Get(key interface{}) (interface{}, bool) {
	n := m.Find(pair{key, nil}, m.Comparer)
	if n == nil {
		return nil, false
	}
	return n.Value.(pair).value, true
}

func (m Map) Set(key, value interface{}) Map {
	node := &treap.Node{pair{key, value}, rand.Intn(1000000), nil, nil}
	n := m.Node.Union(node, m.Comparer, true)
	return Map{n, m.Comparer}
}

func (m Map) Delete(key interface{}) Map {
	n := m.Node.Delete(pair{key, nil}, m.Comparer)
	return Map{n, m.Comparer}
}

func (m Map) ForEach(fn func(key, value interface{})) {
	m.Node.ForEach(func(v interface{}) {
		p := v.(pair)
		fn(p.key, p.value)
	})
}

func (m Map) Count() int {
	result := 0
	m.Node.ForEach(func(_ interface{}) {
		result++
	})
	return result
}

func Example_orderedMap() {
	rand.Seed(42)

	m := NewMap(IntComparer{})
	fmt.Println("Count:", m.Count())

	m = m.Set(52, "hello")
	m = m.Set(53, "world")
	m = m.Set(52, "Hello")

	m.ForEach(func(k, v interface{}) {
		fmt.Println("[", k, "] =", v)
	})
	fmt.Println("Count:", m.Count())

	old := m.Set(500, 500)
	m = m.Delete(53)

	fmt.Println(m.Get(53))
	fmt.Println(old.Get(53))
	fmt.Println(old.Get(52))
	fmt.Println(old.Get(500))

	// Output:
	// Count: 0
	// [ 52 ] = Hello
	// [ 53 ] = world
	// Count: 2
	// <nil> false
	// world true
	// Hello true
	// 500 true
}
