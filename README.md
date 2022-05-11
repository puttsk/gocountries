[![Github](https://img.shields.io/badge/sources-github-green.svg)](https://github.com/puttsk/gocountries/)
 [![GitHub forks](https://img.shields.io/github/stars/puttsk/gocountries.svg?style=social&label=Star)](https://github.com/puttsk/gocountries)

# gocountries

Go module for country information based on ISO-3166-1. The module is built using a code generator from [codegen](https://github.com/puttsk/codegen) project and a data from [world_countries](https://github.com/stefangabos/world_countries) dataset.

## Installation

``` bash
go get github.com/puttsk/gocountries
```

## Build

To build your own module, you need [codegen](https://github.com/puttsk/codegen) and [world_countries](https://github.com/stefangabos/world_countries) data. 

1. Clone `codegen` to the base directory

    ``` bash
    git clone https://github.com/puttsk/codegen.git
    ```

2. Clone `world_countries` dataset to `/data` directory

    ``` bash
    mkdir -p data
    git clone https://github.com/stefangabos/world_countries
    ```

3. Call `make` command to build the module. Make sure that `PYTHON`, `DATA`, and `GENERATOR` in [Makefile](Makefile) use correct version of Python.

    ``` bash
    make
    ```

    This command call `codegen` to generate the code from the dataset and templates file. If success, you should see an output like this.

    ```bash
    python3 ./codegen/generator.py --data data/world_countries/data/countries/_combined/countries.json -t templates/country.go.template -o country.go
    python3 ./codegen/generator.py --data data/world_countries/data/countries/_combined/countries.json -t templates/country_test.go.template -o country_test.go
    go test ./...
    ok      github.com/puttsk/go-countries  (cached)
    ```

## Contributing

TBD

## Authors

* **Putt Sakdhnagool** - *Initial work*

See also the list of [contributors](https://github.com/puttsk/codegen/graphs/contributors) who participated in this project.

## Issues / Feature request

You can submit bug / issues / feature request using [Tracker](https://github.com/puttsk/codegen/issues).

## License

[MIT License](LICENSE)
