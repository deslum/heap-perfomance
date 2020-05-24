# heap-perfomance
benchmarks heap perfomance

```bash
goos: linux
goarch: amd64
pkg: test-heap
BenchmarkMinElementWithTuneHeap-4   	1000000000	         0.000003 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinElementWithSort-4       	1000000000	         0.182 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinElementWithLoop-4       	1000000000	         0.0444 ns/op	       0 B/op	       0 allocs/op
PASS
```
