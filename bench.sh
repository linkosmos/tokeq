echo "" >> BENCHMARK.md
echo "`date` (`go version`)" >> BENCHMARK.md
go test -run=10000 -bench=. -benchmem >> BENCHMARK.md
