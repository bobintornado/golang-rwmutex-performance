```
go test -bench=. -cpu=1,2,3,4 ./performance.go ./performance_test.go
goos: darwin
goarch: amd64
BenchmarkRCache     	  200000	      7665 ns/op
BenchmarkRCache-2   	 1000000	      4176 ns/op
BenchmarkRCache-3   	 2000000	      1380 ns/op
BenchmarkRCache-4   	 3000000	      8339 ns/op
BenchmarkLCache     	  200000	      8309 ns/op
BenchmarkLCache-2   	 1000000	      1977 ns/op
BenchmarkLCache-3   	 3000000	      2596 ns/op
BenchmarkLCache-4   	 3000000	       672 ns/op
PASS
ok  	command-line-arguments	84.080s
```
