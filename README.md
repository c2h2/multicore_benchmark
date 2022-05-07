# multicore_benchmark
Simple math arithmetics, Parallelly benchmarking core speed with GO 

# Build: on mac/linux
```go build multicore_bench.go``` or ```./build_all_platforms.sh ```

# Run: with 4 cores
```./multicore_bench -c 4```
 
# Running results: on a Macbook Air M1
```
$> build/multicore_bench-darwin-arm64 -c 4
hello, benchmarking FP + INT with 4 goroutine(s)
goroutine 0 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 3.0130305s
goroutine 3 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 3.014720292s
goroutine 1 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 3.015916417s
goroutine 2 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 3.019769584s
goroutine 2 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.967434792s
goroutine 3 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.97426s
goroutine 0 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.977723125s
goroutine 1 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.978366667s
goroutine 3 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.385311375s
goroutine 0 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.386120333s
goroutine 2 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.386440166s
goroutine 1 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.393041708s
goroutine 2 finished. Int64, Value = 2176680827608690 Total Time took: 2.393171459s
goroutine 0 finished. Int64, Value = 2176680827608690 Total Time took: 2.395501666s
goroutine 1 finished. Int64, Value = 2176680827608690 Total Time took: 2.39824575s
goroutine 3 finished. Int64, Value = 2176680827608690 Total Time took: 2.3982815s
 ```
 