JWTの署名と署名の検証のベンチマーク。

```
$ go test -bench . -benchmem

goos: darwin
goarch: amd64
pkg: jwt_benchmark
BenchmarkHS256-4                   19519             60527 ns/op           58090 B/op         53 allocs/op
BenchmarkParseHS256-4              12484             94600 ns/op           67416 B/op         86 allocs/op
BenchmarkHS384-4                   22525             53571 ns/op           58442 B/op         53 allocs/op
BenchmarkParseHS384-4              13784             87321 ns/op           67720 B/op         84 allocs/op
BenchmarkHS512-4                   22381             53422 ns/op           58776 B/op         53 allocs/op
BenchmarkParseHS512-4              13635             88083 ns/op           67896 B/op         86 allocs/op
BenchmarkRS256_4096-4                134           9035198 ns/op          137049 B/op        171 allocs/op
BenchmarkParseRS256_4096-4          4021            275272 ns/op           83380 B/op         92 allocs/op
BenchmarkRS384_4096-4                134           8928326 ns/op          137023 B/op        171 allocs/op
BenchmarkParseRS384_4096-4          4125            266985 ns/op           83496 B/op         92 allocs/op
BenchmarkRS512_4096-4                134           8945186 ns/op          137240 B/op        172 allocs/op
BenchmarkParseRS512_4096-4          4476            261421 ns/op           83507 B/op         92 allocs/op
BenchmarkRS256_2048-4                741           1616531 ns/op           89830 B/op        157 allocs/op
BenchmarkParseRS256_2048-4          7179            164180 ns/op           72842 B/op         91 allocs/op
BenchmarkRS384_2048-4                732           1652130 ns/op           90085 B/op        157 allocs/op
BenchmarkParseRS384_2048-4          6613            172359 ns/op           72952 B/op         91 allocs/op
BenchmarkRS512_2048-4                718           1692410 ns/op           89917 B/op        156 allocs/op
BenchmarkParseRS512_2048-4          7982            153963 ns/op           72969 B/op         91 allocs/op
BenchmarkRS256_1024-4               2884            395054 ns/op           74329 B/op        150 allocs/op
BenchmarkParseRS256_1024-4          8118            124688 ns/op           69909 B/op         91 allocs/op
BenchmarkRS384_1024-4               2817            410153 ns/op           74425 B/op        150 allocs/op
BenchmarkParseRS384_1024-4          8930            116380 ns/op           70021 B/op         91 allocs/op
BenchmarkRS512_1024-4               2898            387372 ns/op           74459 B/op        150 allocs/op
BenchmarkParseRS512_1024-4          9009            113898 ns/op           70037 B/op         91 allocs/op
```
