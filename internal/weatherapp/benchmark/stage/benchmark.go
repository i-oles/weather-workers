package stage

import (
	"fmt"
	"math"
	"time"

	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp/runner/stage"
	"main.go/internal/weatherapp/writer"
)

type Benchmark struct {
	stageName string
	writer    writer.PerformanceTestWriter
}

func NewBenchmark(
	stageName string,
	writer writer.PerformanceTestWriter,
) *Benchmark {
	return &Benchmark{
		stageName: stageName,
		writer:    writer,
	}
}

func (b Benchmark) ProcessPerformanceStagesTest(
	runner stage.Runner,
	times int,
) error {
	stageDurations := make([]time.Duration, times)

	for i := 0; i < times; i++ {
		logrus.Infof("%d. processing %s...", i+1, b.stageName)

		startTime := time.Now()

		err := runner.Run()
		if err != nil {
			return fmt.Errorf("failed to run stage: %w", err)
		}

		stageDuration := time.Since(startTime)
		stageDurations[i] = stageDuration

		logrus.Debugf("process no. %v, %s duration: %v\n", i+1, b.stageName, stageDuration)

		err = b.writer.Write(fmt.Sprintf("%v\n", stageDuration))
		if err != nil {
			return fmt.Errorf("failed to write stage duration: %w", err)
		}
	}

	err := b.writer.Write(
		fmt.Sprintf("average: %v\n", calculateAverageStageDuration(stageDurations)),
	)

	if err != nil {
		return fmt.Errorf("failed to write stage duration: %w", err)
	}

	err = b.writer.Write(
		fmt.Sprintf("standard_deviation: %v\n", calculateStandardDeviation(stageDurations)),
	)
	if err != nil {
		return fmt.Errorf("failed to write stage duration: %w", err)
	}

	return nil
}

func calculateAverageStageDuration(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	var sum time.Duration

	for _, duration := range durations {
		sum += duration
	}

	return sum / time.Duration(len(durations))
}

func calculateStandardDeviation(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	avg := calculateAverageStageDuration(durations)

	var stdDeviation float64
	for _, duration := range durations {
		stdDeviation += math.Pow(float64(duration)-float64(avg), 2)
	}

	stdDeviation = math.Sqrt(stdDeviation / float64(len(durations)))

	return time.Duration(stdDeviation)
}
