.PHONY: all

all: test bench

test:
	go test

bench:
	go test -bench Benchmark -run xx -count 5 -benchmem

profile:
	go test -bench BenchmarkMergeSort -run xx -cpuprofile cpu.out -memprofile mem.out

mytest:
	go test -bench . -run xxx -cpuprofile=cpu.out
