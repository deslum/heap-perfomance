# heap-perfomance
benchmarks heap perfomance


GOROOT=/home/yuri/go #gosetup
GOPATH=/home/yuri:/home/yuri/work #gosetup
/home/yuri/go/bin/go test -c -o /tmp/___gobench_test_heap test-heap #gosetup
/tmp/___gobench_test_heap -test.v -test.bench . -test.run ^$ #gosetup
goos: linux
goarch: amd64
pkg: test-heap
BenchmarkMinElementWithTuneHeap-4   	1000000000	         0.000003 ns/op	     0 B/op	       0 allocs/op
BenchmarkMinElementWithSort-4       	1000000000	         0.182 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinElementWithLoop-4       	1000000000	         0.0444 ns/op	       0 B/op	       0 allocs/op
PASS

Process finished with exit code 0
