PYTHON=python3
GENERATOR=./codegen/generator.py
DATA=data/world_countries/data/countries/_combined/countries.json

.DEFAULT: all

.PHONY: add
all: module test

.PHONY: module
module:
	$(PYTHON) $(GENERATOR) --data $(DATA) -t templates/country.go.template -o country.go

.PHONY: test
test:
	$(PYTHON) $(GENERATOR) --data $(DATA) -t templates/country_test.go.template -o country_test.go
	go test -v ./...

.PHONY: profile
profile:
	go clean -testcache
	go test -bench=. -v -memprofile mem.out ./...
	go test -bench=. -v -cpuprofile cpu.out ./...

.PHONY: benchmark
benchmark:
	go test -bench=. -benchmem 