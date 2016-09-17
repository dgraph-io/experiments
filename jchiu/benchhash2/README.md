# Sample results

Note that `ShardedGoMap64` means we have 64 shards of standard Go maps.

Here is what each setup is trying to measure.

* BenchmarkRead: One concurrent read.
* BenchmarkWrite: One concurrent write.
* BenchmarkReadWrite: One concurrent write, followed by one concurrent read.
* BenchmarkRead3Write1: Three reads, one write.
* BenchmarkRead1Write3: Three writes, one read.

```
testing: warning: no tests to run
BenchmarkRead/GoMap-8    	10000000	       131 ns/op
BenchmarkRead/GotomicMap-8         	10000000	       161 ns/op
BenchmarkRead/ShardedGoMap8-8      	10000000	       125 ns/op
BenchmarkRead/ShardedGoMap16-8     	10000000	       125 ns/op
BenchmarkRead/ShardedGoMap32-8     	10000000	       124 ns/op
BenchmarkRead/ShardedGoMap64-8     	10000000	       125 ns/op

BenchmarkWrite/GoMap-8             	 5000000	       319 ns/op
BenchmarkWrite/GotomicMap-8        	 2000000	       848 ns/op
BenchmarkWrite/ShardedGoMap8-8     	 5000000	       297 ns/op
BenchmarkWrite/ShardedGoMap16-8    	 5000000	       289 ns/op
BenchmarkWrite/ShardedGoMap32-8    	 5000000	       284 ns/op
BenchmarkWrite/ShardedGoMap64-8    	 5000000	       281 ns/op

BenchmarkReadWrite/GoMap-8         	 3000000	       471 ns/op
BenchmarkReadWrite/GotomicMap-8    	 1000000	      1117 ns/op
BenchmarkReadWrite/ShardedGoMap8-8 	 3000000	       494 ns/op
BenchmarkReadWrite/ShardedGoMap16-8         	 3000000	       472 ns/op
BenchmarkReadWrite/ShardedGoMap32-8         	 3000000	       449 ns/op
BenchmarkReadWrite/ShardedGoMap64-8         	 3000000	       442 ns/op

BenchmarkRead3Write1/GoMap-8                	 2000000	       879 ns/op
BenchmarkRead3Write1/GotomicMap-8           	 1000000	      1885 ns/op
BenchmarkRead3Write1/ShardedGoMap8-8        	 2000000	       888 ns/op
BenchmarkRead3Write1/ShardedGoMap16-8       	 2000000	       813 ns/op
BenchmarkRead3Write1/ShardedGoMap32-8       	 2000000	       780 ns/op
BenchmarkRead3Write1/ShardedGoMap64-8       	 2000000	       763 ns/op

BenchmarkRead1Write3/GoMap-8                	 1000000	      1108 ns/op
BenchmarkRead1Write3/GotomicMap-8           	  500000	      3085 ns/op
BenchmarkRead1Write3/ShardedGoMap8-8        	 1000000	      1113 ns/op
BenchmarkRead1Write3/ShardedGoMap16-8       	 1000000	      1060 ns/op
BenchmarkRead1Write3/ShardedGoMap32-8       	 1000000	      1029 ns/op
BenchmarkRead1Write3/ShardedGoMap64-8       	 1000000	      1005 ns/op
PASS
ok  	github.com/dgraph-io/experiments/jchiu/benchhash2	51.936s
```