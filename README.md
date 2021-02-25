# go-test-example
This repository is for testing golang.

## Usage
Run test
```
$ go test -v
=== RUN   TestAdd
=== RUN   TestAdd/1+2
=== RUN   TestAdd/2+5
=== RUN   TestAdd/99+1
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/1+2 (0.00s)
    --- PASS: TestAdd/2+5 (0.00s)
    --- PASS: TestAdd/99+1 (0.00s)
PASS
```

Run benchmark
```
goos: darwin
goarch: amd64
BenchmarkSliceAppend-8                   9890283               127 ns/op              42 B/op          0 allocs/op
BenchmarkSliceAppendAlocOnce-8          463262860                3.89 ns/op            0 B/op          0 allocs/op
PASS
```
