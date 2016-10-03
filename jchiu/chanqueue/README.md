We have seen numerous articles complaining about how slow channels are. However,
in the following benchmarks, channels seem to do pretty well. Here are the
results.

```
BenchmarkChan-8     	20000000	        70.3 ns/op
BenchmarkQueue-8    	10000000	       222 ns/op
BenchmarkCQueue-8   	 1000000	      2495 ns/op
```

`Queue` uses a simple Go array that can grow.

`CQueue` uses a circular buffer. If you try to push to a full buffer, it blocks. If
you try to pop from an empty buffer, it blocks. This behaves just like a buffered
channels.

The tests are constructed such that there is one goroutine pushing stuff into a
queue or channel. Once it is done, it marks the queue as done. The main thread
keeps popping until the queue is empty and the queue is marked as done.

Channel seems to run much faster than `Queue` or `CQueue`.