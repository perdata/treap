// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Package treap implements a persistent treap (tree/heap combination).
//
// https://en.wikipedia.org/wiki/Treap
//
// A treap is a binary search tree for storing ordered distinct values
// (duplicates not allowed). In addition, each node actually a random
// priority field which is stored in heap order (i.e. all children
// have lower priority than the parent)
//
// This provides the basis for efficient immutable ordered Set
// operations.  See the ordered map example for how this can be used
// as an ordered map
//
// Much of this is based on "Fast Set Operations Using Treaps"
// by Guy E Blelloch and Margaret Reid-Miller:
// https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
//
// Benchmark
//
//    $ go test --bench=. -benchmem
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/perdata/treap
//    BenchmarkUnion-4             	    1000	   1456648 ns/op	  939846 B/op	   19580 allocs/op
//    BenchmarkIntersection-4      	     500	   3112224 ns/op	 1719838 B/op	   35836 allocs/op
//     BenchmarkIntersectionMap-4   	    1000	   1354798 ns/op	  364991 B/op	      84 allocs/op
//    PASS
//
package treap

// Comparer compares two values. The return value is zero if the
// values are equal, negative if the first is smaller and positive
// otherwise.
type Comparer interface {
	Compare(left, right interface{}) int
}

// Node is the basic recursive treap data structure
type Node struct {
	Value       interface{}
	Priority    int
	Left, Right *Node
}

// ForEach does inorder traversal of the treap
func (n *Node) ForEach(fn func(v interface{})) {
	if n != nil {
		n.Left.ForEach(fn)
		fn(n.Value)
		n.Right.ForEach(fn)
	}
}

// Find finds the node in the treap with matching value
func (n *Node) Find(v interface{}, c Comparer) *Node {
	for {
		if n == nil {
			return nil
		}
		diff := c.Compare(n.Value, v)
		switch {
		case diff == 0:
			return n
		case diff < 0:
			n = n.Right
		case diff > 0:
			n = n.Left
		}
	}
}

// Union combines any two treaps. In case of duplicates, the overwrite
// field controls whether the union keeps the original value or
// whether it is updated based on value in the "other" arg
func (n *Node) Union(other *Node, c Comparer, overwrite bool) *Node {
	if n == nil {
		return other
	}
	if other == nil {
		return n
	}

	if n.Priority < other.Priority {
		other, n, overwrite = n, other, !overwrite
	}

	left, dupe, right := other.Split(n.Value, c)
	value := n.Value
	if overwrite && dupe != nil {
		value = dupe.Value
	}
	left = n.Left.Union(left, c, overwrite)
	right = n.Right.Union(right, c, overwrite)
	return &Node{value, n.Priority, left, right}
}

// Split splits the treap into all nodes that compare less-than, equal
// and greater-than the provided value.  The resulting values are
// properly formed treaps or nil if they contain no values.
func (n *Node) Split(v interface{}, c Comparer) (left, mid, right *Node) {
	leftp, rightp := &left, &right
	for {
		if n == nil {
			*leftp = nil
			*rightp = nil
			return left, nil, right
		}

		root := &Node{n.Value, n.Priority, nil, nil}
		diff := c.Compare(n.Value, v)
		switch {
		case diff < 0:
			*leftp = root
			root.Left = n.Left
			leftp = &root.Right
			n = n.Right
		case diff > 0:
			*rightp = root
			root.Right = n.Right
			rightp = &root.Left
			n = n.Left
		default:
			*leftp = n.Left
			*rightp = n.Right
			return left, root, right
		}
	}
}

// Intersection returns a new treap with all the common values in the
// two treaps.
//
// see https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
// "Fast Set Operations Using Treaps"
//   by Guy E Blelloch and Margaret Reid-Miller.
//
// The algorithm is a very slight variation on that.
func (n *Node) Intersection(other *Node, c Comparer) *Node {
	if n == nil || other == nil {
		return nil
	}

	if n.Priority < other.Priority {
		n, other = other, n
	}

	left, found, right := other.Split(n.Value, c)
	left = n.Left.Intersection(left, c)
	right = n.Right.Intersection(right, c)

	if found == nil {
		// TODO: use a destructive join as both left/right are copies
		return left.join(right)
	}

	return &Node{n.Value, n.Priority, left, right}
}

// Delete removes a node if it exists.
func (n *Node) Delete(v interface{}, c Comparer) *Node {
	left, _, right := n.Split(v, c)
	return left.join(right)
}

// see https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
// "Fast Set Operations Using Treaps"
//   by Guy E Blelloch and Margaret Reid-Miller.
//
// The algorithm is a very slight variation on that provided there.
//
// Note that all nodes in n have priority <= that of "other" for
// this call to work correctly.  It traverses  the right spine of n
// and left-spine of other, merging things along the way
//
// The algorithm is not that  different from zipping up a spine
func (n *Node) join(other *Node) *Node {
	var result *Node
	resultp := &result
	for {
		if n == nil {
			*resultp = other
			return result
		}
		if other == nil {
			*resultp = n
			return result
		}

		if n.Priority <= other.Priority {
			root := &Node{n.Value, n.Priority, n.Left, nil}
			*resultp = root
			resultp = &root.Right
			n = n.Right
		} else {
			root := &Node{other.Value, other.Priority, nil, other.Right}
			*resultp = root
			resultp = &root.Left
			other = other.Left
		}
	}
}
