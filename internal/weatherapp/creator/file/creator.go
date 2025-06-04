package file

import (
	"fmt"
	"os"
	"path/filepath"

	"main.go/internal/weatherapp/configuration"
	"main.go/internal/weatherapp/creator"
)

const (
	jsonExtension     = ".json"
	csvExtension      = ".csv"
	resultFilePrefix  = "results"
	testingFilePrefix = "testing"
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
	filePath := c.buildFilePath()

	if err := c.validateIfFileExists(filePath); err != nil {
		return nil, fmt.Errorf("failed validating file: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed creating file: %w", err)
	}

	return file, nil
}

func (c *Creator) buildFilePath() string {
	return filepath.Join(c.cfg.FilesDirName, c.buildFileName())
}

func (c *Creator) buildFileName() string {
	if c.cfg.PerformanceTest {
		return fmt.Sprintf("%s_%s_%dtimes%s",
			testingFilePrefix,
			c.cfg.Mode,
			c.cfg.ExecutionRepeatCount,
			csvExtension,
		)
	}

	return fmt.Sprintf("%s_%s%s", resultFilePrefix, c.cfg.Mode, jsonExtension)
}

func (c *Creator) validateIfFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", filePath)
	}

	return nil
}
