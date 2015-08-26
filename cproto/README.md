# Benchmark of various encoding systems.
## Generated via https://github.com/cloudflare/goser

### Populate
BenchmarkPopulatePb	20000000	      1264 ns/op	     417 B/op	      16 allocs/op
BenchmarkPopulateGogopb	50000000	       370 ns/op	      48 B/op	       3 allocs/op
BenchmarkPopulateCapnp	10000000	      2714 ns/op	     114 B/op	       2 allocs/op

### Marshal
BenchmarkMarshalJSON	 2000000	      9690 ns/op	  61.71 MB/s	     756 B/op	      17 allocs/op
BenchmarkMarshalPb	20000000	      1092 ns/op	 264.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkMarshalGogopb	50000000	       607 ns/op	 475.69 MB/s	     320 B/op	       1 allocs/op
BenchmarkMarshalCapnp	100000000	       167 ns/op	2769.93 MB/s	       8 B/op	       0 allocs/op

### Unmarshal
BenchmarkUnmarshalJSON	  500000	     28562 ns/op	  20.94 MB/s	    2275 B/op	      26 allocs/op
BenchmarkUnmarshalPb	 5000000	      3377 ns/op	  85.57 MB/s	     871 B/op	      20 allocs/op
BenchmarkUnmarshalGogopb	20000000	      1136 ns/op	 254.27 MB/s	     266 B/op	       7 allocs/op
BenchmarkUnmarshalCapnp	20000000	       936 ns/op	 495.28 MB/s	     263 B/op	       5 allocs/op
BenchmarkUnmarshalCapnpZeroCopy	50000000	       366 ns/op	1267.18 MB/s	      91 B/op	       3 allocs/op
go tool pprof --svg goser.test cpu.prof > cpu.svg

