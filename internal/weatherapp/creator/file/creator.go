package file

import (
	"fmt"
	"os"
	"path/filepath"

	"main.go/internal/weatherapp/configuration"
	"main.go/internal/weatherapp/creator"
)

const (
	extJSON      = ".json"
	extCSV       = ".csv"
	prefixResult = "result"
	prefixTest   = "test"
)

var _ creator.FileCreator = (*Creator)(nil)

type Creator struct {
	cfg configuration.Configuration
}

func NewCreator(cfg configuration.Configuration) *Creator {
	return &Creator{
		cfg: cfg,
	}
}

func (c *Creator) Create() (*os.File, error) {
	fileName := fmt.Sprintf("%s_%s%s", prefixResult, c.cfg.Mode, extJSON)

	maxWorkingProducers := ""
	if c.cfg.Mode == "mode_5" {
		maxWorkingProducers = fmt.Sprintf("_%dmax_working_producers", c.cfg.MaxWorkingProducers)
	}

	if c.cfg.PerformanceTest {
		fileName = fmt.Sprintf("%s_%s__%dtimes_%dproducers_%dconsumers%s%s",
			prefixTest,
			c.cfg.Mode,
			c.cfg.ExecutionRepeatCount,
			c.cfg.ProducerNumber,
			c.cfg.ConsumerNumber,
			maxWorkingProducers,
			extCSV,
		)
	}

	filePath := filepath.Join(c.cfg.FilesDirName, fileName)

	if err := c.ensureFileNotExists(filePath); err != nil {
		return nil, fmt.Errorf("failed validating file: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed creating file: %w", err)
	}

	return file, nil
}

func (c *Creator) ensureFileNotExists(filePath string) error {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", filePath)
	}

	return nil
}
