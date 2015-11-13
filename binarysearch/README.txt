Based on benchmarks, Iterative implementation runs ~25% faster than recursive.

$ go test -v -bench .
testing: warning: no tests to run
PASS
BenchmarkRec_100-6   	10000000	       139 ns/op
BenchmarkRec_1000-6  	10000000	       179 ns/op
BenchmarkRec_10000-6 	10000000	       220 ns/op
BenchmarkIter_100-6  	20000000	       110 ns/op
BenchmarkIter_1000-6 	10000000	       138 ns/op
BenchmarkIter_10000-6	10000000	       166 ns/op
ok  	github.com/dgraph-io/experiments/binarysearch	11.642s

