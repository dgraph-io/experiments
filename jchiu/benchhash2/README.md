Note that `ShardedGoMap64` means we have 64 shards of standard Go maps.

Here is what each setup is trying to measure.

* Control2: Two `rand.Uint32` calls.
* Control5: Five `rand.Uint32` calls.
* BenchmarkRead: One concurrent read. Subtract timings by Control.
* BenchmarkWrite: One concurrent write. Subtract timings by Control2.
* BenchmarkReadWrite: One concurrent write, followed by one concurrent read. Subtract timings by Control3.
* BenchmarkRead3Write1: Three reads, one write. Subtract timings by Control5.
* BenchmarkRead1Write3: Three writes, one read. Subtract timings by Control7.

To run the test, do

```
$ go test -bench=.
```

# Sample results

In theory, you should subtract the time taken to compute the random integers.
For example, for `BenchmarkRead`, GoMap takes 131-59.4ns while GotomicMap takes
163-59.4ns, which is about 45% slower.

```
testing: warning: no tests to run
BenchmarkControl-8       	30000000	        59.4 ns/op
BenchmarkControl2-8      	20000000	       116 ns/op
BenchmarkControl3-8      	10000000	       174 ns/op
BenchmarkControl5-8      	 5000000	       289 ns/op
BenchmarkControl7-8      	 3000000	       403 ns/op
BenchmarkRead/GoMap-8    	10000000	       131 ns/op
BenchmarkRead/GotomicMap-8         	10000000	       163 ns/op
BenchmarkRead/ShardedGoMap8-8      	10000000	       125 ns/op
BenchmarkRead/ShardedGoMap16-8     	10000000	       126 ns/op
BenchmarkRead/ShardedGoMap32-8     	10000000	       124 ns/op
BenchmarkRead/ShardedGoMap64-8     	10000000	       124 ns/op
BenchmarkWrite/GoMap-8             	 5000000	       320 ns/op
BenchmarkWrite/GotomicMap-8        	 2000000	       867 ns/op
BenchmarkWrite/ShardedGoMap8-8     	 5000000	       300 ns/op
BenchmarkWrite/ShardedGoMap16-8    	 5000000	       291 ns/op
BenchmarkWrite/ShardedGoMap32-8    	 5000000	       286 ns/op
BenchmarkWrite/ShardedGoMap64-8    	 5000000	       283 ns/op
BenchmarkReadWrite/GoMap-8         	 3000000	       473 ns/op
BenchmarkReadWrite/GotomicMap-8    	 1000000	      1153 ns/op
BenchmarkReadWrite/ShardedGoMap8-8 	 3000000	       495 ns/op
BenchmarkReadWrite/ShardedGoMap16-8         	 3000000	       471 ns/op
BenchmarkReadWrite/ShardedGoMap32-8         	 3000000	       454 ns/op
BenchmarkReadWrite/ShardedGoMap64-8         	 3000000	       443 ns/op
BenchmarkRead3Write1/GoMap-8                	 2000000	       857 ns/op
BenchmarkRead3Write1/GotomicMap-8           	 1000000	      1859 ns/op
BenchmarkRead3Write1/ShardedGoMap8-8        	 2000000	       878 ns/op
BenchmarkRead3Write1/ShardedGoMap16-8       	 2000000	       816 ns/op
BenchmarkRead3Write1/ShardedGoMap32-8       	 2000000	       780 ns/op
BenchmarkRead3Write1/ShardedGoMap64-8       	 2000000	       768 ns/op
BenchmarkRead1Write3/GoMap-8                	 1000000	      1086 ns/op
BenchmarkRead1Write3/GotomicMap-8           	  500000	      3072 ns/op
BenchmarkRead1Write3/ShardedGoMap8-8        	 1000000	      1135 ns/op
BenchmarkRead1Write3/ShardedGoMap16-8       	 1000000	      1062 ns/op
BenchmarkRead1Write3/ShardedGoMap32-8       	 1000000	      1030 ns/op
BenchmarkRead1Write3/ShardedGoMap64-8       	 1000000	      1014 ns/op
PASS
ok  	github.com/dgraph-io/experiments/jchiu/benchhash2	61.849s
```