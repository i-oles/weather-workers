# ABOUT Weather Workers App

Aplikacja producer-consumer analizująca pogodę z podanego okresu.

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

program zapisze żądane wyniki w `/assets`.  

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

### performance tests:

Ponieważ aplikacja ma różne tryby - możemy odpalić testy performatywne, które zmierzą nam: 

run-performance-test:
run-race-performance-test:
go run cmd/weatherapp/main.go
go run -race cmd/weatherapp/main.go -profile="testing"

### unittests:
run:  
`go test -v ./...` or `go test --race -v ./...`  

or using makefile:  
`make test` or `make test-race`

Rzeczy do poruszenia: współdzielona pamięć (mutex), ograniczenia w strzałach do api (semaphore)

Program powinien być w stanie zliczyć czas swojego działania, żeby móc porównać pod względem wydajności kolejne etapy.
Wystarczy zwykłe liczenie czasu z biblioteką time

TODO:

- popraw tresc zapisywanych plikow testu benchmarku
- datarace consumer? producer? (mode_3,4,5)j
- dodaj flage na odapalnie testow z iloscia powtórzen execution
- opisz w readme testing -> zwykle odpalenie testow i opisz testy benchmarkowe
- w nazwie plikow result i test powinno byc widoczne ile bylo consumerow i producerow
- 