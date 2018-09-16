# treap

[![Status](https://travis-ci.com/perdata/treap.svg?branch=master)](https://travis-ci.com/perdata/treap?branch=master)
[![GoDoc](https://godoc.org/github.com/perdata/treap?status.svg)](https://godoc.org/github.com/perdata/treap)
[![codecov](https://codecov.io/gh/perdata/treap/branch/master/graph/badge.svg)](https://codecov.io/gh/perdata/treap)
[![GoReportCard](https://goreportcard.com/badge/github.com/perdata/treap)](https://goreportcard.com/report/github.com/perdata/treap)

Package treap implements a persistent sorted set datastructure using a combination tree/heap or [treap](https://en.wikipedia.org/wiki/Treap).

The algorithms are mostly based on [Fast Set Operations Using Treaps](https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf)

Although the package is oriented towards ordered sets, it is simple to convert it to work as a persistent map.


##  Benchmark stats

```sh
$ go test --bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/perdata/treap
BenchmarkUnion-4             	    1000	   1456648 ns/op	  939846 B/op	   19580 allocs/op
BenchmarkIntersection-4      	     500	   3112224 ns/op	 1719838 B/op	   35836 allocs/op
BenchmarkIntersectionMap-4   	    1000	   1354798 ns/op	  364991 B/op	      84 allocs/op
PASS
```