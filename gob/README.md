Benchmark was to test how much decline in performance we'd get for Gob.Encode, v/s directly writing our binary data to io.Writer.

These are the results:

```
$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkGobEncode_1K	  200000	     11388 ns/op
BenchmarkGobEncode_1M	    5000	    253726 ns/op
BenchmarkGobEncode_50M	     100	  18707543 ns/op
BenchmarkEncode_1K	100000000	        10.0 ns/op
BenchmarkEncode_1M	200000000	         9.84 ns/op
BenchmarkEncode_50M	200000000	         9.91 ns/op
ok  	github.com/dgraph-io/experiments/gob	52.960s

```
