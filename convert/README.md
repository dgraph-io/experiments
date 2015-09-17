### About

Tested CRC with ISO and ECMA polynomials, and a custom library implementing
murmur hash. None of them showed any collissions for 100 million unique ids.

### Test Results

```
# go test -v .
=== RUN TestUseCrc
--- PASS: TestUseCrc (0.00s)
=== RUN TestUseCrc_ISOCollissions
--- PASS: TestUseCrc_ISOCollissions (11.55s)
=== RUN TestUseCrc_ECMACollissions
--- PASS: TestUseCrc_ECMACollissions (11.07s)
=== RUN TestUseMurmur_Collissions
--- PASS: TestUseMurmur_Collissions (11.28s)
PASS
ok  	github.com/dgraph-io/experiments/convert	34.015s
```

### Benchmark Results

```
BenchmarkUseCrc_ISO	10000000	       168 ns/op
BenchmarkUseCrc_ECMA	10000000	       168 ns/op
BenchmarkUseMurmur	10000000	       194 ns/op
```
