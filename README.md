# ABOUT 

Aplikacja Weather Workers App producer-consumer analizująca pogodę z podanego okresu.

Program na podstawie danych o polskich miastach uderza do api `http://Open-Meteo.com`
i zapisuje do pliku następujące dane:

1. miasto z największą srednią temperaturą
2. miasto gdzie najczęściej występuje mgła (weather code 45)
3. miasto gdzie jest najczęściej czyste niebo (weather code 0)

Program wejsciowo dostaje plik `assets/pl172.json` z informacjami o miastach z bazy danych z `http://Simplemaps.com`

# DEVELOPMENT

# RUNNING

You can run this app in few modes:  
`mode_1` --> 1 producer, 1 consumer (singe thread)  
`mode_2` --> 1 producer, 1 consumer (goroutines)  
`mode_3` --> 1 producer, n consumers  
`mode_4` --> n producers, n consumers  
`mode_5` --> n producers, n consumers, but only k producers can work in the same time.

open config file `config/dev.json`

```
   "AnalysisDurationInMonths": 6,           --> modify time period for analysis in months
   "Mode": "mode_2",                        --> specify given mode
   "ConsumerNumber": 5,                     --> specify number of consumers (modify only if you use mode_3, mode_4, mode_5)
   "ProducerNumber": 5,                     --> specify number of producers (modify only if you use mode_4, mode_5)
   "MaxWorkingProducers": 5                 --> specify number of working producers in the same time (only if you use mode_5)
```

run:  
`go run cmd/weatherapp/main.go` or `go run -race cmd/weatherapp/main.go`

or using makefile:  
`make run` or `make run-race`

program zapisze żądane wyniki w `/assets`   

przykladowy response:
```
{
  "highest_avg_temp": {
    "city_name": "Katowice",
    "value": 5.5729735883424425
  },
  "hours_with_fog": {
    "city_name": "Lębork",
    "value": 83
  },
  "hours_with_full_sun": {
    "city_name": "Braniewo",
    "value": 3166
  }
}
```
# TESTING

### PERFORMANCE TESTS:

Ponieważ aplikacja ma różne tryby - możemy odpalić testy performatywne,  
w których możemy okreslic ile razy chcemy odpalić ten sam proces (aplikację w konkretnym mode).
Dzięki tym testom możemy sprawdzić wydajnosc poszczególnych modeów.
Po odpaleniu zostaną zapisane do pliku poszczególne czasy każdej egzekucji, sredni czas egzekucji, a także czas standard deviation.
Przykładowy zapis:

```
execution_1: 4.454012416s
execution_2: 4.447440375s
execution_3: 4.452014083s
average_execution: 4.451155624s
standard_deviation: 2.750835ms
```

open config file `config/testing.json`  

specify settings in config (look up RUNNING section)  
dodatkowo:

```
  "MockApi": true,                  --> żeby nie przeciążyć API `http://Open-Meteo.com`, mockujemy sredni czas jednego strzału do API
  "ExecutionRepeatCount": 3,        --> specify how many times you want to execute aplication
  "PerformanceTest": true,          --> set to true for running performance tests 
```

run:  
`go run cmd/weatherapp/main.go -profile="testing"` or `go run -race cmd/weatherapp/main.go -profile="testing"`  

or using makefile:  
`make run-performance-test` or `make run-race-performance-test`

program zapisze żądane wyniki w `/assets`  

### UNIT TESTS:
run:  
`go test -v ./...` or `go test --race -v ./...`  

or using makefile:  
`make test` or `make test-race`

- datarace consumer? producer? (mode_3,4,5)j
- w nazwie plikow result i test powinno byc widoczne ile bylo consumerow i producerow