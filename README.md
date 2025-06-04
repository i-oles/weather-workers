# ABOUT Weather Workers App

Aplikacja producer-consumer analizująca pogodę z ostatniego półrocza.

Program na podstawie danych o polskich miastach uderza do api `http://Open-Meteo.com`
i zapisuje do pliku następujące dane:
1. miasto z największą srednią temperaturą
2. miasto gdzie najczęściej występuje mgła (weather code 45)
3. miasto gdzie jest najczęściej czyste niebo (weather code 0)

Program jako input dostaje plik json z informacjami o miastach z bazy danych z `http://Simplemaps.com`
Ścieżka do pliku to: `assets/pl172.json`

Aplikację można odpalić w kliku trybach:

1. 1 producer, 1 consumer (singe thread)
2. 1 producer, 1 consumer (goroutines)
3. 1 producer, n consumers
4. n producers, n consumers
5. n producers, n consumers, but only k producers can work in the same time.

Wyniki zwrócone przez program znajdziesz w `assets`.
Pattern nazewnictwa zwróconych plików:

`results_mode<number_of_mode>.json`

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

# DEVELOPMENT




# RUNNING


# TESTING

go run cmd/weatherapp/main.go

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