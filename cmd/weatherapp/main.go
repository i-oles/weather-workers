package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/aggregator/weather"
	"main.go/internal/weatherapp/benchmark/stage"
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
	stageRunner "main.go/internal/weatherapp/runner/stage"
	"main.go/internal/weatherapp/runner/stage/stagefour"
	"main.go/internal/weatherapp/runner/stage/stageone"
	"main.go/internal/weatherapp/runner/stage/stagethree"
	"main.go/internal/weatherapp/runner/stage/stagetwo"
	"main.go/internal/weatherapp/saver/memory"
	"main.go/internal/weatherapp/writer"
	"main.go/internal/weatherapp/writer/benchmark"
	logWriter "main.go/internal/weatherapp/writer/log"
	"main.go/internal/weatherapp/writer/result"
	"main.go/pkg/config"
)

func main() {
	startTime := time.Now()

	cfg := configuration.Configuration{}

	err := config.GetConfig("./config", &cfg)
	if err != nil {
		logrus.Fatalf("error during getting config: %v", err)
	}

	log.Println(cfg.Pretty())

	var apiOpenMeteo openmeteo.API

	apiOpenMeteo = api.NewOpenMeteo(cfg.APIURL, cfg.AnalysisDurationInMonths)
	if cfg.MockAPI {
		apiOpenMeteo = mock.NewOpenMeteo()
	}

	var producer weatherAppProducer.Producer

	producer = memoryProducer.New(apiOpenMeteo)

	if cfg.Stage == "stage5" {
		semaphore := semaphore.NewWeighted(int64(cfg.MaxWorkingProducers))
		producer = semaphoredecorator.New(producer, semaphore)
	}

	if cfg.LogProducedMsg {
		producer = logProducer.New(producer)
	}

	var preSavers []presaver.WeatherAppPreSaver
	preSavers = []presaver.WeatherAppPreSaver{
		&avgtemp.PreSaver{},
		&sunnyhours.PreSaver{},
		&foggyhours.PreSaver{},
	}

	if cfg.Stage != "stage1" && cfg.Stage != "stage2" {
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
		logrus.Fatalf(
			"error during reading source file %v: %v", cfg.SourceFileName, err)
	}

	shortCitiesInfo := make([]weatherapp.ShortCityInfo, len(citiesInfo))
	for i, cityInfo := range citiesInfo {
		shortCitiesInfo[i] = weatherapp.GetShortCityInfo(cityInfo)
	}

	resultsFileCreator := file.NewCreator(cfg)

	resultsFile, err := resultsFileCreator.Create()
	if err != nil {
		logrus.Fatalf("error during creating results file: %v", err)
	}

	defer resultsFile.Close()

	switch cfg.Stage {
	case "stage1":
		runner := stageone.NewRunner(producer, consumer, shortCitiesInfo)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Stage, cfg.StageRepeatCount)
			if err != nil {
				logrus.Fatalf("error performance tests for %s: %v", cfg.Stage, err)
			}

			return
		}

		err = runner.Run()
		if err != nil {
			logrus.Fatalf("error during running stage1: %v", err)
		}
	case "stage2":
		runner := stagetwo.NewRunner(producer, consumer, shortCitiesInfo)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Stage, cfg.StageRepeatCount)
			if err != nil {
				logrus.Fatalf("error performance tests for %s: %v", cfg.Stage, err)
			}

			return
		}

		err = runner.Run()
		if err != nil {
			logrus.Fatalf("error during running stage2: %v", err)
		}
	case "stage3":
		runner := stagethree.NewRunner(producer, consumer, shortCitiesInfo, cfg.ConsumerNumber)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Stage, cfg.StageRepeatCount)
			if err != nil {
				logrus.Fatalf("error performance tests for %s: %v", cfg.Stage, err)
			}

			return
		}

		err = runner.Run()
		if err != nil {
			logrus.Fatalf("error during running stage3: %v", err)
		}
	case "stage4", "stage5":
		runner := stagefour.NewRunner(
			producer,
			consumer,
			shortCitiesInfo,
			cfg.ConsumerNumber,
			cfg.ProducerNumber,
		)
		if cfg.PerformanceTest {
			err = processPerformanceTests(runner, resultsFile, cfg.Stage, cfg.StageRepeatCount)
			if err != nil {
				logrus.Fatalf("error performance tests for %s: %v", cfg.Stage, err)
			}

			return
		}

		err = runner.Run()
		if err != nil {
			logrus.Fatalf("error during running %v: %v", cfg.Stage, err)
		}
	case "stage6":

	default:
		logrus.Fatalf("Unknown stage: %v", cfg.Stage)
	}

	results := saver.GetResults()

	var resultWriter writer.WeatherAppResultWriter

	resultWriter = result.NewWriter(resultsFile)
	if cfg.LogResults {
		resultWriter = logWriter.New(resultWriter)
	}

	err = resultWriter.Write(results)
	if err != nil {
		logrus.Fatalf("error during writing the result file: %v", err)
	}

	logrus.Infof("Process duration: %v", time.Since(startTime))
}

func processPerformanceTests(
	runner stageRunner.Runner,
	file *os.File,
	stageName string,
	stageRepeatCount int,
) error {
	stageDurationWriter := benchmark.NewWriter(file)
	stageBenchmark := stage.NewBenchmark(stageName, stageDurationWriter)

	err := stageBenchmark.ProcessPerformanceStagesTest(runner, stageRepeatCount)
	if err != nil {
		return fmt.Errorf("error during performance tests for %s: %w", stageName, err)
	}

	return nil
}
