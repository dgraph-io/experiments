# Benchmarks

## BoltDB

Without copying the resulting byte slice from Bolt. **Unsafe**
```
$ go test -bench BenchmarkRead .
testing: warning: no tests to run
PASS
BenchmarkReadBolt_1024	  500000	      3858 ns/op
BenchmarkReadBolt_10KB	  500000	      3738 ns/op
BenchmarkReadBolt_500KB	 1000000	      3141 ns/op
BenchmarkReadBolt_1MB	 1000000	      3026 ns/op
ok  	github.com/dgraph-io/experiments/db	102.513s
```

Copying the resulting byte slice. **Safe**
```
$ go test -bench BenchmarkRead .
testing: warning: no tests to run
PASS
BenchmarkReadBolt_1024	  200000	      6760 ns/op
BenchmarkReadBolt_10KB	  100000	     21249 ns/op
BenchmarkReadBolt_500KB	   10000	    214449 ns/op
BenchmarkReadBolt_1MB	    3000	    350712 ns/op
ok  	github.com/dgraph-io/experiments/db	80.890s
```

## RocksDB

```
$ go test -bench BenchmarkGet .
PASS
BenchmarkGet_valsize1024	  300000	      5715 ns/op
BenchmarkGet_valsize10KB	   50000	     27619 ns/op
BenchmarkGet_valsize500KB	    2000	    604185 ns/op
BenchmarkGet_valsize1MB	    2000	   1064685 ns/op
ok  	github.com/dgraph-io/dgraph/store	55.029s
```
