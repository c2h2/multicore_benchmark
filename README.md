[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/c2h2/multicore_benchmark/) 
[![license](https://img.shields.io/github/license/c2h2/multicore_benchmark.svg?style=plastic)]() 
[![download_count](https://img.shields.io/github/downloads/c2h2/multicore_benchmark/total.svg?style=plastic)](https://github.com/c2h2/multicore_benchmark/releases) 
[![download](https://img.shields.io/github/release/c2h2/multicore_benchmark.svg?style=plastic)](https://github.com/c2h2/multicore_benchmark/releases)

# multicore_benchmark
Simple math arithmetics, Parallelly benchmarking core speed with GO 


# Build: on mac/linux
```go build multicore_bench.go``` or ```./build_all_platforms.sh ```

# Run: with 4 cores
```./multicore_bench -c 4```
 
# Running results:

Macbook Air M1
```
$> multicore_bench_build/multicore_bench-darwin-arm64 -c 4
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
 
AMD Ryzen 5950x
 ```
$> multicore_bench_build/multicore_bench-linux-amd64 -c 4
hello, benchmarking FP + INT with 4 goroutine(s)
goroutine 1 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 2.238745459s
goroutine 2 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 2.241890144s
goroutine 0 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 2.271218761s
goroutine 3 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 2.271348542s
goroutine 1 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.879430684s
goroutine 3 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.879786115s
goroutine 2 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.880157397s
goroutine 0 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 2.881713884s
goroutine 1 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.433871896s
goroutine 2 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.434115687s
goroutine 0 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.434667099s
goroutine 3 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 2.435343671s
goroutine 3 finished. Int64, Value = 2176680827608690 Total Time took: 1.923438975s
goroutine 1 finished. Int64, Value = 2176680827608690 Total Time took: 1.924397299s
goroutine 0 finished. Int64, Value = 2176680827608690 Total Time took: 1.927111899s
goroutine 2 finished. Int64, Value = 2176680827608690 Total Time took: 1.930444902s
```

Rockchip RK3568
```
$> multicore_bench_build/multicore_bench-linux-arm64 -c 4
hello, benchmarking FP + INT with 4 goroutine(s)
goroutine 1 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 43.942383374s
goroutine 3 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 44.361646957s
goroutine 2 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 44.40480342s
goroutine 0 finished. Float32, Value(overflowed) = 7.0368744e+13 Total Time took: 44.698104908s
goroutine 0 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 1m0.227978635s
goroutine 3 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 1m0.436453531s
goroutine 2 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 1m0.512283955s
goroutine 1 finished. Float64, Value = 2.176682290419792e+15 Total Time took: 1m1.046790884s
goroutine 0 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 26.505736288s
goroutine 2 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 26.668196097s
goroutine 1 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 26.758366313s
goroutine 3 finished. Int32, Value(overflowed) = -1090425452 Total Time took: 26.901846494s
goroutine 0 finished. Int64, Value = 2176680827608690 Total Time took: 30.168077448s
goroutine 2 finished. Int64, Value = 2176680827608690 Total Time took: 30.252918332s
goroutine 1 finished. Int64, Value = 2176680827608690 Total Time took: 30.265622458s
goroutine 3 finished. Int64, Value = 2176680827608690 Total Time took: 30.45509206s
```
