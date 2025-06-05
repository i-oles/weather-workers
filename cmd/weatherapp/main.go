package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"golang.org/x/sync/semaphore"
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/aggregator/weather"
	"main.go/internal/weatherapp/benchmark/execution"
	"main.go/internal/weatherapp/config"
	"main.go/internal/weatherapp/configuration"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	logConsumer "main.go/internal/weatherapp/consumer/log"
	memoryConsumer "main.go/internal/weatherapp/consumer/memory"
	"main.go/internal/weatherapp/creator/file"
	"main.go/internal/weatherapp/openmeteo"
	"main.go/internal/weatherapp/openmeteo/api"
	"main.go/internal/weatherapp/openmeteo/mock"
	"main.go/internal/weatherapp/presaver"
	"main.go/internal/weatherapp/presaver/avgtemp"
	"main.go/internal/weatherapp/presaver/foggyhours"
	"main.go/internal/weatherapp/presaver/mutexdecorator"
	"main.go/internal/weatherapp/presaver/sunnyhours"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	logProducer "main.go/internal/weatherapp/producer/log"
	memoryProducer "main.go/internal/weatherapp/producer/memory"
	"main.go/internal/weatherapp/producer/semaphoredecorator"
	fileReader "main.go/internal/weatherapp/reader/file"
	modeRunner "main.go/internal/weatherapp/runner/mode"
	modeFour "main.go/internal/weatherapp/runner/mode/modefour"
	modeOne "main.go/internal/weatherapp/runner/mode/modeone"
	modeThree "main.go/internal/weatherapp/runner/mode/modethree"
	modeTwo "main.go/internal/weatherapp/runner/mode/modetwo"
	"main.go/internal/weatherapp/saver/memory"
	"main.go/internal/weatherapp/writer"
	"main.go/internal/weatherapp/writer/benchmark"
	logWriter "main.go/internal/weatherapp/writer/log"
	"main.go/internal/weatherapp/writer/result"
)

func main() {
	cfg := configuration.Configuration{}

	err := config.GetConfig("./config", &cfg)
	if err != nil {
		slog.Error("error during getting config", slog.Any("error", err))
	}

	log.Println(cfg.Pretty())

	err = run(cfg)
	if err != nil {
		slog.Error("error running app", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(cfg configuration.Configuration) error {
	var apiOpenMeteo openmeteo.API

	apiOpenMeteo = api.NewOpenMeteo(cfg.APIURL, cfg.AnalysisDurationInMonths)
	if cfg.MockAPI {
		apiOpenMeteo = mock.NewOpenMeteo()
	}

	var producer weatherAppProducer.Producer

	producer = memoryProducer.New(apiOpenMeteo)

	if cfg.LogProducedMsg {
		producer = logProducer.New(producer)
	}

	var preSavers []presaver.WeatherAppPreSaver
	preSavers = []presaver.WeatherAppPreSaver{
		&avgtemp.PreSaver{},
		&sunnyhours.PreSaver{},
		&foggyhours.PreSaver{},
	}

	if cfg.Mode != "mode_1" && cfg.Mode != "mode_2" {
		preSavers = []presaver.WeatherAppPreSaver{
			mutexdecorator.New(&avgtemp.PreSaver{}),
			mutexdecorator.New(&sunnyhours.PreSaver{}),
			mutexdecorator.New(&foggyhours.PreSaver{}),
		}
	}

	saver := memory.NewSaver(preSavers)
	aggregator := weather.NewAggregator()

	var consumer weatherAppConsumer.Consumer

	consumer = memoryConsumer.New(aggregator, saver)
	if cfg.LogConsumedMsg {
		consumer = logConsumer.New(consumer)
	}

	reader := fileReader.New()
	sourcePathFile := filepath.Join(cfg.FilesDirName, cfg.SourceFileName)

	citiesInfo, err := reader.Read(sourcePathFile)
	if err != nil {
		return fmt.Errorf(
			"error during reading source file %v: %w", cfg.SourceFileName, err)
	}

	shortCitiesInfo := make([]weatherapp.ShortCityInfo, len(citiesInfo))
	for i, cityInfo := range citiesInfo {
		shortCitiesInfo[i] = weatherapp.GetShortCityInfo(cityInfo)
	}

	resultsFileCreator := file.NewCreator(cfg)

	resultsFile, err := resultsFileCreator.Create()
	if err != nil {
		return fmt.Errorf("error during creating results file: %w", err)
	}

	defer resultsFile.Close()

	switch cfg.Mode {
	case "mode_1":
		runner := modeOne.NewRunner(producer, consumer, shortCitiesInfo)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Mode, cfg.ExecutionRepeatCount)
			if err != nil {
				return fmt.Errorf("error performance tests for %s: %w", cfg.Mode, err)
			}

			return nil
		}

		err = runner.Run()
		if err != nil {
			return fmt.Errorf("error during running %s: %w", cfg.Mode, err)
		}
	case "mode_2":
		runner := modeTwo.NewRunner(producer, consumer, shortCitiesInfo)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Mode, cfg.ExecutionRepeatCount)
			if err != nil {
				return fmt.Errorf("error performance tests for %s: %w", cfg.Mode, err)
			}

			return nil
		}

		err = runner.Run()
		if err != nil {
			return fmt.Errorf("error during running %s: %w", cfg.Mode, err)
		}
	case "mode_3":
		runner := modeThree.NewRunner(producer, consumer, shortCitiesInfo, cfg.ConsumerNumber)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Mode, cfg.ExecutionRepeatCount)
			if err != nil {
				return fmt.Errorf("error performance tests for %s: %w", cfg.Mode, err)
			}

			return nil
		}

		err = runner.Run()
		if err != nil {
			return fmt.Errorf("error during running %s: %w", cfg.Mode, err)
		}
	case "mode_4", "mode_5":
		if cfg.Mode == "mode_5" {
			sem := semaphore.NewWeighted(int64(cfg.MaxWorkingProducers))
			producer = semaphoredecorator.New(producer, sem)
		}

		runner := modeFour.NewRunner(
			producer,
			consumer,
			shortCitiesInfo,
			cfg.ConsumerNumber,
			cfg.ProducerNumber,
		)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Mode, cfg.ExecutionRepeatCount)
			if err != nil {
				return fmt.Errorf("error performance tests for %s: %w", cfg.Mode, err)
			}

			return nil
		}

		err = runner.Run()
		if err != nil {
			return fmt.Errorf("error during running %s: %w", cfg.Mode, err)
		}
	default:
		return fmt.Errorf("unknown execution: %v", cfg.Mode)
	}

	results := saver.GetResults()

	var resultWriter writer.WeatherAppResultWriter

	resultWriter = result.NewWriter(resultsFile)
	if cfg.LogResults {
		resultWriter = logWriter.New(resultWriter)
	}

	err = resultWriter.Write(results)
	if err != nil {
		return fmt.Errorf("error during writing the result file: %w", err)
	}

	return nil
}

func processPerformanceTests(
	runner modeRunner.Runner,
	file *os.File,
	mode string,
	executionRepeatCount int,
) error {
	modeDurationWriter := benchmark.NewWriter(file)
	modeBenchmark := execution.NewBenchmark(mode, modeDurationWriter)

	err := modeBenchmark.ProcessExecutionPerformanceTest(runner, executionRepeatCount)
	if err != nil {
		return fmt.Errorf("error during execution performance tests for %s: %w", mode, err)
	}

	return nil
}
