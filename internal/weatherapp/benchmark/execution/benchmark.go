package execution

import (
	"fmt"
	"log/slog"
	"math"
	"time"

	"main.go/internal/weatherapp/runner/mode"
	"main.go/internal/weatherapp/writer"
)

type Benchmark struct {
	mode   string
	writer writer.PerformanceTestWriter
}

func NewBenchmark(
	mode string,
	writer writer.PerformanceTestWriter,
) *Benchmark {
	return &Benchmark{
		mode:   mode,
		writer: writer,
	}
}

func (b Benchmark) ProcessExecutionPerformanceTest(
	runner mode.Runner,
	times int,
) error {
	executionDurations := make([]time.Duration, times)

	for i := 0; i < times; i++ {
		slog.Info("Processing...",
			slog.Int("execution_number", i+1),
			slog.String("mode", b.mode),
		)

		startTime := time.Now()

		err := runner.Run()
		if err != nil {
			return fmt.Errorf("failed to run program with execution %s: %w", b.mode, err)
		}

		executionDuration := time.Since(startTime)
		executionDurations[i] = executionDuration

		err = b.writer.Write(fmt.Sprintf("execution_%d: %v\n", i+1, executionDuration))
		if err != nil {
			return fmt.Errorf("failed to write execution duration: %w", err)
		}

		slog.Info("execution finished:",
			slog.Int("execution_number", i+1),
			slog.String("mode", b.mode),
			slog.Duration("duration", executionDuration),
		)
	}

	err := b.writer.Write(
		fmt.Sprintf("average_execution: %v\n", calculateAverageExecutionDuration(executionDurations)),
	)

	if err != nil {
		return fmt.Errorf("failed to write execution duration: %w", err)
	}

	err = b.writer.Write(
		fmt.Sprintf("standard_deviation: %v\n", calculateStandardDeviation(executionDurations)),
	)
	if err != nil {
		return fmt.Errorf("failed to write execution duration: %w", err)
	}

	slog.Info("performance test done")

	return nil
}

func calculateAverageExecutionDuration(durations []time.Duration) time.Duration {
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

	avg := calculateAverageExecutionDuration(durations)

	var stdDeviation float64
	for _, duration := range durations {
		stdDeviation += math.Pow(float64(duration)-float64(avg), 2)
	}

	stdDeviation = math.Sqrt(stdDeviation / float64(len(durations)))

	return time.Duration(stdDeviation)
}
