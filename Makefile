.PHONY: bench-cpu bench-mem

bench-cpu:
	@cd "$(filter-out $@,$(MAKECMDGOALS))" && \
	go test -bench . -benchmem -o main.test -run ^$$ -cpuprofile=cpu.prof -memprofile=mem.prof && \
	go tool pprof -http=:8080 ./main.test cpu.prof

bench-mem:
	@cd "$(filter-out $@,$(MAKECMDGOALS))" && \
	go test -bench . -benchmem -o main.test -run ^$$ -cpuprofile=cpu.prof -memprofile=mem.prof && \
	go tool pprof -http=:8080 ./main.test mem.prof