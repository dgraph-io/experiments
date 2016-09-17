# General setup
There are three different setups `Read`, `Write`, `ReadWrite`. Here are the main parameters.

* `b.N`: This is the total number of reps for go bench to do its measurement.
* `n`: For each rep, we will read/write `n` elements to **one single hash** across multiple goroutines.
* `q`: For each rep, we will use `q` goroutines.

Each of the three setups has the same structure below:

* Assume `n` is divisible by `q`.
* Create an array `work` of size `n`. This is a list of key-value pairs.
* Start timer.
* For each of the `b.N` reps, do the following:
  * Create an empty hash map `h`.
  * Create `q` goroutines.
	* Each goroutine will either read or write to `h`.
	* Each goroutine works on an equal part of the array `work`.

Note that we do not use channels at all as that will cause some unnecessary blocking.

For the `ReadWrite` setup, we ask the user for an additional parameter `fracRead`. There will be `fracRead * q` goroutines which do reads and `(1-fracRead) * q` goroutines which do writes. To be clear, each goroutine scans `n/q` elements.

We did not use `RunParallel` because it doesn't fit the structure of our test setup.

Please see `benchhash.go` for the test setups.

# Results

Run `run.sh` to run the benchmarks. The results are found in the `results` subdirectory.

Here is the main line in `run.sh`:

```
go test -cpu $NUMCPU -benchn 100000 -benchq $q -bench=. > results/benchhash.q$q.txt
```

You might want to want to tweak the parameter `NUMCPU` in `run.sh`.

## Interpreting results

Note that `BenchReadWrite7` means `7/10` of the goroutines are doing reads. Generally, reads are cheaper, so `BenchReadWrite9` should take less time than `BenchReadWrite1`.

Note that `ShardedGoMap16` means we use 16 shards of GoMaps.

Here are some results for `q=1000`.

```
n=100000 q=1000
BenchmarkRead/GoMap-2   	     200	   6882485 ns/op
BenchmarkRead/GotomicMap-2         	     100	  20416675 ns/op
BenchmarkRead/ShardedGoMap4-2      	     200	   6269036 ns/op
BenchmarkRead/ShardedGoMap8-2      	     300	   5677571 ns/op
BenchmarkRead/ShardedGoMap16-2     	     300	   5333423 ns/op
BenchmarkRead/ShardedGoMap32-2     	     300	   5153572 ns/op
BenchmarkWrite/GoMap-2             	      50	  20158550 ns/op
BenchmarkWrite/GotomicMap-2        	      10	 102329390 ns/op
BenchmarkWrite/ShardedGoMap4-2     	      50	  20343076 ns/op
BenchmarkWrite/ShardedGoMap8-2     	     100	  16701933 ns/op
BenchmarkWrite/ShardedGoMap16-2    	     100	  15416442 ns/op
BenchmarkWrite/ShardedGoMap32-2    	     100	  14433085 ns/op
BenchmarkReadWrite1/GoMap-2        	     100	  17983497 ns/op
BenchmarkReadWrite1/GotomicMap-2   	      20	 101842924 ns/op
BenchmarkReadWrite1/ShardedGoMap4-2         	      50	  20708505 ns/op
BenchmarkReadWrite1/ShardedGoMap8-2         	     100	  17904104 ns/op
BenchmarkReadWrite1/ShardedGoMap16-2        	     100	  15886365 ns/op
BenchmarkReadWrite1/ShardedGoMap32-2        	     100	  14662240 ns/op
BenchmarkReadWrite3/GoMap-2                 	     100	  17667777 ns/op
BenchmarkReadWrite3/GotomicMap-2            	      20	  81755781 ns/op
BenchmarkReadWrite3/ShardedGoMap4-2         	     100	  19808166 ns/op
BenchmarkReadWrite3/ShardedGoMap8-2         	     100	  17530502 ns/op
BenchmarkReadWrite3/ShardedGoMap16-2        	     100	  15399216 ns/op
BenchmarkReadWrite3/ShardedGoMap32-2        	     100	  13931323 ns/op
BenchmarkReadWrite5/GoMap-2                 	     100	  15603580 ns/op
BenchmarkReadWrite5/GotomicMap-2            	      30	  59283460 ns/op
BenchmarkReadWrite5/ShardedGoMap4-2         	     100	  17145842 ns/op
BenchmarkReadWrite5/ShardedGoMap8-2         	     100	  15334110 ns/op
BenchmarkReadWrite5/ShardedGoMap16-2        	     100	  13167841 ns/op
BenchmarkReadWrite5/ShardedGoMap32-2        	     100	  11573469 ns/op
BenchmarkReadWrite7/GoMap-2                 	     100	  15413127 ns/op
BenchmarkReadWrite7/GotomicMap-2            	      30	  45866904 ns/op
BenchmarkReadWrite7/ShardedGoMap4-2         	     100	  15388264 ns/op
BenchmarkReadWrite7/ShardedGoMap8-2         	     100	  13555623 ns/op
BenchmarkReadWrite7/ShardedGoMap16-2        	     100	  11592148 ns/op
BenchmarkReadWrite7/ShardedGoMap32-2        	     100	  10214670 ns/op
BenchmarkReadWrite9/GoMap-2                 	     100	  13929774 ns/op
BenchmarkReadWrite9/GotomicMap-2            	      50	  28248601 ns/op
BenchmarkReadWrite9/ShardedGoMap4-2         	     100	  11725283 ns/op
BenchmarkReadWrite9/ShardedGoMap8-2         	     200	   9930883 ns/op
BenchmarkReadWrite9/ShardedGoMap16-2        	     200	   8302751 ns/op
BenchmarkReadWrite9/ShardedGoMap32-2        	     200	   7418499 ns/op
PASS
ok  	github.com/jchiu0/experimental/benchhash	71.801s
```