# multicore_benchmark
Simple math arithmetics, Parallelly benchmarking core speed with GO 

# build
```go build multicore_bench.go```

# run with 4 cores
```./multicore_bench -c 4```
 
# demo on M1 mac
```
./multicore_bench -c 4
hello, benchmarking FP + INT with 4 goroutine(s)
goroutine 2 finished. Float64, Value = 5.807400525456593e+16 Total Time took: 15.488039417s
goroutine 0 finished. Float64, Value = 5.807400525456593e+16 Total Time took: 15.525536542s
goroutine 3 finished. Float64, Value = 5.807400525456593e+16 Total Time took: 15.548415167s
goroutine 1 finished. Float64, Value = 5.807400525456593e+16 Total Time took: 15.549120792s
goroutine 1 finished. Int64, Value = 58073997698937742 Total Time took: 12.642762916s
goroutine 2 finished. Int64, Value = 58073997698937742 Total Time took: 12.648014333s
goroutine 3 finished. Int64, Value = 58073997698937742 Total Time took: 12.657186458s
goroutine 0 finished. Int64, Value = 58073997698937742 Total Time took: 12.657670375s
 ```
