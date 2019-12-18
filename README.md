# treap

[![Status](https://travis-ci.com/perdata/treap.svg?branch=master)](https://travis-ci.com/perdata/treap?branch=master)
[![GoDoc](https://godoc.org/github.com/perdata/treap?status.svg)](https://godoc.org/github.com/perdata/treap)
[![codecov](https://codecov.io/gh/perdata/treap/branch/master/graph/badge.svg)](https://codecov.io/gh/perdata/treap)
[![GoReportCard](https://goreportcard.com/badge/github.com/perdata/treap)](https://goreportcard.com/report/github.com/perdata/treap)

Package treap implements a [persistent](https://en.wikipedia.org/wiki/Persistent_data_structure) sorted set datastructure using a combination tree/heap or [treap](https://en.wikipedia.org/wiki/Treap).

The algorithms are mostly based on [Fast Set Operations Using Treaps](https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf)

Although the package is oriented towards ordered sets, it is simple to convert it to work as a persistent map.  There is a working [example](https://godoc.org/github.com/perdata/treap#example-package--OrderedMap) showing how to do this.


##  Benchmark stats

The most interesting benchmark is the performance of insert where a
single random key is inserted into a 5k sized map.  As the example
shows, the treap structure does well here as opposed to a regular
persistent map (which involves full copying).  This benchmark does not
take into account the fact that the regular maps are not sorted unlike
treaps. 

The intersection benchmark compares the case where two 10k sets with
5k in common being interesected. The regular persistent array is about
30% faster but this is still respectable showing for treaps. 


```sh
$ go test --bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/perdata/treap
BenchmarkInsert-4                   	 1000000	      2347 ns/op	    1719 B/op	      36 allocs/op
BenchmarkInsertRegularMap-4         	    2000	    890745 ns/op	  336311 B/op	       8 allocs/op
BenchmarkIntersection-4             	     500	   3125772 ns/op	 1719838 B/op	   35836 allocs/op
BenchmarkIntersectionRegularMap-4   	     500	   2436519 ns/op	  718142 B/op	     123 allocs/op
BenchmarkUnion-4                    	    1000	   1451047 ns/op	  939846 B/op	   19580 allocs/op
BenchmarkDiff-4                     	     500	   3280823 ns/op	 1742080 B/op	   36298 allocs/op
PASS
```
