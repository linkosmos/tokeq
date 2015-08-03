echo "" >> BENCHMARK.md
echo `date` >> BENCHMARK.md
go test -run=10000 -bench=. -benchmem >> BENCHMARK.md
