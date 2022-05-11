GENERATOR=python3 ./codegen/generator.py
DATA=data/world_countries/data/countries/_combined/countries.json

all: module test

module:
	$(GENERATOR) --data $(DATA) -t templates/go.template -o country.go

test:
	$(GENERATOR) --data $(DATA) -t templates/go_test.template -o country_test.go
	go test ./...
	go test -bench=.