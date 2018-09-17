// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package treap_test

import (
	"github.com/perdata/treap"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestUnion(t *testing.T) {
	rand.Seed(42)

	input := []int{}
	for kk := 0; kk < 5000; kk++ {
		input = append(input, rand.Intn(10000000))
	}

	part1 := ToTreap(input[:len(input)/2])
	part2 := ToTreap(input[len(input)/2:])
	combined := part1.Union(part2, IntComparer{}, false)
	all := ToTreap(input)

	if !reflect.DeepEqual(ToArray(combined), ToArray(all)) {
		t.Fatal("Union diverged")
	}
}

func TestIntersection(t *testing.T) {
	rand.Seed(42)

	input1 := []int{}
	input2 := []int{}
	input3 := []int{}
	start := 0
	for kk := 0; kk < 5000; kk++ {
		start += 1 + rand.Intn(100)
		input1 = append(input1, start)
		start += 1 + rand.Intn(100)
		input2 = append(input2, start)
		start += 1 + rand.Intn(100)
		input3 = append(input3, start)
	}

	input1 = append(input1, input3...)
	input2 = append(input2, input3...)
	rand.Shuffle(len(input1), func(i, j int) {
		input1[i], input1[j] = input1[j], input1[i]
	})
	rand.Shuffle(len(input2), func(i, j int) {
		input2[i], input2[j] = input2[j], input2[i]
	})

	set1, set2, common := ToTreap(input1), ToTreap(input2), ToTreap(input3)
	intersection := set1.Intersection(set2, IntComparer{})

	if !reflect.DeepEqual(ToArray(intersection), ToArray(common)) {
		t.Fatal("Intersection diverged")
	}
}

func TestDiff(t *testing.T) {
	rand.Seed(42)

	input1 := []int{}
	input2 := []int{}
	input3 := []int{}
	start := 0
	for kk := 0; kk < 5000; kk++ {
		start += 1 + rand.Intn(100)
		input1 = append(input1, start)
		start += 1 + rand.Intn(100)
		input2 = append(input2, start)
		start += 1 + rand.Intn(100)
		input3 = append(input3, start)
	}

	input13 := append(input1[:5000:5000], input3...)
	input23 := append(input2[:5000:5000], input3...)
	rand.Shuffle(len(input13), func(i, j int) {
		input13[i], input13[j] = input13[j], input13[i]
	})
	rand.Shuffle(len(input23), func(i, j int) {
		input23[i], input23[j] = input23[j], input23[i]
	})

	set1, set2 := ToTreap(input13), ToTreap(input23)
	diff := set1.Diff(set2, IntComparer{})

	if !reflect.DeepEqual(ToArray(diff), ToArray(ToTreap(input1))) {
		t.Fatal("Diff diverged")
	}
}

func BenchmarkInsert(b *testing.B) {
	rand.Seed(42)

	input := []int{}
	for kk := 0; kk < 10000; kk++ {
		input = append(input, rand.Intn(10000000))
	}

	x := ToTreap(input)
	c := IntComparer{}
	for kk := 0; kk < b.N; kk++ {
		insert := ToTreap([]int{rand.Intn(10000000)})
		x.Union(insert, c, false)
	}
}

func BenchmarkInsertRegularMap(b *testing.B) {
	rand.Seed(42)

	input := map[interface{}]bool{}
	for kk := 0; kk < 10000; kk++ {
		input[rand.Intn(10000000)] = true
	}

	for kk := 0; kk < b.N; kk++ {
		insert := rand.Intn(10000000)
		clone := make(map[interface{}]bool, len(input))
		for key, value := range input {
			clone[key] = value
		}
		clone[insert] = true
	}
}

func BenchmarkIntersection(b *testing.B) {
	rand.Seed(42)

	input1 := []int{}
	input2 := []int{}
	input3 := []int{}
	start := 0
	for kk := 0; kk < 5000; kk++ {
		start += 1 + rand.Intn(100)
		input1 = append(input1, start)
		start += 1 + rand.Intn(100)
		input2 = append(input2, start)
		start += 1 + rand.Intn(100)
		input3 = append(input3, start)
	}

	input1 = append(input1, input3...)
	input2 = append(input2, input3...)
	rand.Shuffle(len(input1), func(i, j int) {
		input1[i], input1[j] = input1[j], input1[i]
	})
	rand.Shuffle(len(input2), func(i, j int) {
		input2[i], input2[j] = input2[j], input2[i]
	})

	set1, set2 := ToTreap(input1), ToTreap(input2)

	for kk := 0; kk < b.N; kk++ {
		set1.Intersection(set2, IntComparer{})
	}
}

func BenchmarkIntersectionRegularMap(b *testing.B) {
	rand.Seed(42)

	input1 := []int{}
	input2 := []int{}
	input3 := []int{}
	start := 0
	for kk := 0; kk < 5000; kk++ {
		start += 1 + rand.Intn(100)
		input1 = append(input1, start)
		start += 1 + rand.Intn(100)
		input2 = append(input2, start)
		start += 1 + rand.Intn(100)
		input3 = append(input3, start)
	}

	input1 = append(input1, input3...)
	input2 = append(input2, input3...)
	rand.Shuffle(len(input1), func(i, j int) {
		input1[i], input1[j] = input1[j], input1[i]
	})
	rand.Shuffle(len(input2), func(i, j int) {
		input2[i], input2[j] = input2[j], input2[i]
	})

	map1, map2 := map[interface{}]bool{}, map[interface{}]bool{}
	for _, elt := range input1 {
		map1[elt] = true
	}
	for _, elt := range input2 {
		map2[elt] = true
	}

	for kk := 0; kk < b.N; kk++ {
		common := map[interface{}]bool{}
		keys := []interface{}{}
		for k, v := range map1 {
			if map2[k] {
				common[k] = v
				keys = append(keys, k)
			}
		}
		sort.Sort(IntArray(keys))
		_ = common
	}
}

func BenchmarkUnion(b *testing.B) {
	rand.Seed(42)

	input := []int{}
	for kk := 0; kk < 10000; kk++ {
		input = append(input, rand.Intn(10000000))
	}

	part1 := ToTreap(input[:len(input)/2])
	part2 := ToTreap(input[len(input)/2:])

	for kk := 0; kk < b.N; kk++ {
		part1.Union(part2, IntComparer{}, false)
	}
}

func BenchmarkDiff(b *testing.B) {
	rand.Seed(42)

	input1 := []int{}
	input2 := []int{}
	input3 := []int{}
	start := 0
	for kk := 0; kk < 5000; kk++ {
		start += 1 + rand.Intn(100)
		input1 = append(input1, start)
		start += 1 + rand.Intn(100)
		input2 = append(input2, start)
		start += 1 + rand.Intn(100)
		input3 = append(input3, start)
	}

	input13 := append(input1[:5000:5000], input3...)
	input23 := append(input2[:5000:5000], input3...)
	rand.Shuffle(len(input13), func(i, j int) {
		input13[i], input13[j] = input13[j], input13[i]
	})
	rand.Shuffle(len(input23), func(i, j int) {
		input23[i], input23[j] = input23[j], input23[i]
	})

	set1, set2 := ToTreap(input13), ToTreap(input23)

	for kk := 0; kk < b.N; kk++ {
		set1.Diff(set2, IntComparer{})
	}
}

func ToTreap(v []int) *treap.Node {
	var t *treap.Node
	for _, elt := range v {
		priority := rand.Intn(10000000)
		t = t.Union(&treap.Node{elt, priority, nil, nil}, IntComparer{}, false)
	}
	return t
}

func ToArray(n *treap.Node) []int {
	if n == nil {
		return nil
	}
	l, r := ToArray(n.Left), ToArray(n.Right)
	l = append([]int(nil), l...)
	return append(append(l, n.Value.(int)), r...)
}

type IntComparer struct{}

func (i IntComparer) Compare(left, right interface{}) int {
	return left.(int) - right.(int)
}

type IntArray []interface{}

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j int) bool {
	return ia[i].(int) < ia[j].(int)
}

func (ia IntArray) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}
