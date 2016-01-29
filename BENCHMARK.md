Mon Aug 3 20:14:52 EEST 2015
PASS
BenchmarkGoQueryFindP                            5000     277185 ns/op     46059 B/op      929 allocs/op
BenchmarkStandardLibraryTokenFindP                500    3698731 ns/op    291695 B/op     9240 allocs/op
BenchmarkStandardLibraryNodeFindP               20000      78340 ns/op        64 B/op        0 allocs/op
BenchmarkDissectNodesFindP                      20000      68640 ns/op        64 B/op        0 allocs/op
ok    github.com/linkosmos/tokeq  8.126s

Mon Nov 23 06:53:16 EET 2015
PASS
BenchmarkGoQueryFindP-8                          5000     210632 ns/op     46068 B/op      928 allocs/op
BenchmarkStandardLibraryTokenFindP-8              500    2657110 ns/op    340032 B/op     8970 allocs/op
BenchmarkStandardLibraryNodeFindP-8             30000      59291 ns/op        44 B/op        0 allocs/op
BenchmarkDissectNodesFindP-8                    30000      56398 ns/op        44 B/op        0 allocs/op
ok    github.com/linkosmos/tokeq  7.356s

Fri Jan 29 19:03:35 EET 2016 (go1.6rc1)
PASS
BenchmarkGoQueryFindP-8                          5000     212107 ns/op     46062 B/op      928 allocs/op
BenchmarkStandardLibraryTokenFindP-8              500    2574948 ns/op    290719 B/op     8971 allocs/op
BenchmarkStandardLibraryNodeFindP-8             30000      56711 ns/op        43 B/op        0 allocs/op
BenchmarkDissectNodesFindP-8                    30000      53779 ns/op        43 B/op        0 allocs/op
ok    github.com/linkosmos/tokeq  7.119s
