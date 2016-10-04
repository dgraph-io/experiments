We have seen numerous articles complaining about how slow channels are. However,
in the following benchmarks, channels seem to do pretty well.

We will test four different implementations: channels, simple array, circular
buffer and finally Gringo, which is a lock-free ring buffer. You can get it by:

```
go get github.com/textnode/gringo/...
```

Here are the results.

```
BenchmarkChan-8     	20000000	        67.0 ns/op
BenchmarkQueue-8    	10000000	       172 ns/op
BenchmarkCQueue-8   	 5000000	       307 ns/op
BenchmarkGringo-8   	20000000	        81.8 ns/op
```

The tests are constructed such that there is one goroutine pushing stuff into a
queue or channel. The main thread keeps popping until the queue is empty.

Channel seems to run faster than the other implementations. This might change if
we have more producers or more consumers.