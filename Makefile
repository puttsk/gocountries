PYTHON=python3
GENERATOR=./codegen/generator.py
DATA=data/world_countries/data/countries/_combined/countries.json

all: module test

module:
	$(PYTHON) $(GENERATOR) --data $(DATA) -t templates/country.go.template -o country.go

test:
	$(PYTHON) $(GENERATOR) --data $(DATA) -t templates/country_test.go.template -o country_test.go
	go test ./...

benchmark:
	go test -bench=.