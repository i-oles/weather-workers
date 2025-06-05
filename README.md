# ABOUT

**Weather Workers App** is a producer-consumer application that analyzes weather data over a specified period of time.

Based on data about Polish cities, the program queries the API `http://Open-Meteo.com` and saves the following results to a file:

- The city with the highest average temperature
- The city with the most frequent fog
- The city with the most frequent clear skies

As input, the program uses the file `assets/pl172.json` containing information about Polish cities from the database `http://Simplemaps.com`.

---

# DEVELOPMENT

install go:  
```
sudo apt install golang-go        →   ubuntu/debian  

brew install go                   →   macOSX  
```

install packages:  

```
go mod install
```

---

# RUNNING

You can run the app in several modes:
- `mode_1` → 1 producer, 1 consumer (single thread)
- `mode_2` → 1 producer, 1 consumer (using goroutines)
- `mode_3` → 1 producer, multiple consumers
- `mode_4` → multiple producers, multiple consumers
- `mode_5` → multiple producers, multiple consumers, but only **k** producers can work concurrently

Open the config file `config/dev.json` and set the following values:

```json
"AnalysisDurationInMonths": 6,          // Duration of analysis in months
"Mode": "mode_2",                       // Mode to run the application
"ConsumerNumber": 5,                    // Number of consumers (used only in mode_3, mode_4, mode_5)
"ProducerNumber": 5,                    // Number of producers (used only in mode_4, mode_5)
"MaxWorkingProducers": 5                // Max number of concurrently working producers (only in mode_5)
```

Run the app with:
```bash
go run cmd/weatherapp/main.go
```

The program will save the results to `/assets`.

Example response:
```json
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

---

# TESTING

## UNIT TESTS:

To run unit tests, use:
```bash
go test -v ./...
```

## PERFORMANCE TESTS:

To check application performance in various modes, you can run performance tests.  
You specify how many times the application should be executed in a given mode.  
After execution, the program will save the execution duration of each run,  
the average execution time, and the standard deviation to a file.

Example output:
```
execution_1: 4.454012416s
execution_2: 4.447440375s
execution_3: 4.452014083s
average_execution: 4.451155624s
standard_deviation: 2.750835ms
```

Open the config file `config/testing.json` and set the following values  
(see the RUNNING section for the shared settings).  
Additionally:

```json
"MockApi": true,                 // Mocks the API call duration to avoid overloading Open-Meteo API
"ExecutionRepeatCount": 3,       // Number of times to execute the application
"PerformanceTest": true          // Enables performance test mode
```

Run with:
```bash
go run cmd/weatherapp/main.go -profile="testing"
```

The program will save the performance results to `/assets`.