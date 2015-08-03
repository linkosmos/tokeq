Mon Aug 3 20:14:52 EEST 2015
PASS
BenchmarkGoQueryFindP                            5000     277185 ns/op     46059 B/op      929 allocs/op
BenchmarkStandardLibraryTokenFindP                500    3698731 ns/op    291695 B/op     9240 allocs/op
BenchmarkStandardLibraryNodeFindP               20000      78340 ns/op        64 B/op        0 allocs/op
BenchmarkDissectNodesFindP                      20000      68640 ns/op        64 B/op        0 allocs/op
ok    github.com/linkosmos/tokeq  8.126s
