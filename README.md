JWTの署名と署名の検証のベンチマーク。

```
$ go test -bench . -benchmem

goos: darwin
goarch: amd64
pkg: jwt_benchmark
BenchmarkHS256-4                   18709             63630 ns/op           58081 B/op         53 allocs/op
BenchmarkParseHS256-4              10000            109524 ns/op           67416 B/op         86 allocs/op
BenchmarkHS384-4                   20810             55230 ns/op           58441 B/op         53 allocs/op
BenchmarkParseHS384-4              13678             88151 ns/op           67720 B/op         84 allocs/op
BenchmarkHS512-4                   22117             54232 ns/op           58786 B/op         53 allocs/op
BenchmarkParseHS512-4              13332             87954 ns/op           67896 B/op         86 allocs/op
BenchmarkRS256_4096-4                124           9420939 ns/op          137299 B/op        171 allocs/op
BenchmarkParseRS256_4096-4          4118            278670 ns/op           83383 B/op         92 allocs/op
BenchmarkRS384_4096-4                134           8990093 ns/op          136481 B/op        171 allocs/op
BenchmarkParseRS384_4096-4          3991            291525 ns/op           83498 B/op         92 allocs/op
BenchmarkRS512_4096-4                130           9247958 ns/op          136933 B/op        171 allocs/op
BenchmarkParseRS512_4096-4          3457            298543 ns/op           83506 B/op         92 allocs/op
BenchmarkRS256_2048-4                711           1738377 ns/op           89846 B/op        157 allocs/op
BenchmarkParseRS256_2048-4          5799            176023 ns/op           72838 B/op         91 allocs/op
BenchmarkRS384_2048-4                716           1727068 ns/op           89898 B/op        156 allocs/op
BenchmarkParseRS384_2048-4          5774            173602 ns/op           72953 B/op         91 allocs/op
BenchmarkRS512_2048-4                668           1715124 ns/op           89816 B/op        156 allocs/op
BenchmarkParseRS512_2048-4          7659            176751 ns/op           72968 B/op         91 allocs/op
```
