package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"main.go/internal/weatherapp/configuration"
	"main.go/internal/weatherapp/creator"
)

var _ creator.FileCreator = (*Creator)(nil)

type Creator struct {
	cfg    configuration.Configuration
	ext    string
	prefix string
}

func NewResultFileCreator(cfg configuration.Configuration) *Creator {
	return &Creator{
		cfg:    cfg,
		ext:    ".json",
		prefix: "result",
	}
}

func NewTestFileCreator(cfg configuration.Configuration) *Creator {
	return &Creator{
		cfg:    cfg,
		ext:    ".csv",
		prefix: "test",
	}
}

func (c *Creator) Create() (*os.File, error) {
	var filename strings.Builder

	filename.WriteString(c.prefix)
	filename.WriteString("_")
	filename.WriteString(c.cfg.Mode)
	filename.WriteString("_")

	if c.prefix == "test" {
		filename.WriteString(fmt.Sprintf("_%dtimes", c.cfg.ExecutionRepeatCount))
	}

	switch c.cfg.Mode {
	case "mode_1", "mode_2":
		filename.WriteString("_1producer_1consumer")
	case "mode_3":
		filename.WriteString(fmt.Sprintf("_1producer_%dconsumer", c.cfg.ConsumerNumber))
	case "mode_4":
		filename.WriteString(fmt.Sprintf("_%dproducer_%dconsumer", c.cfg.ProducerNumber, c.cfg.ConsumerNumber))
	case "mode_5":
		filename.WriteString(fmt.Sprintf("_%dproducer_%dconsumer_%dmax_working_producers",
			c.cfg.ProducerNumber,
			c.cfg.ConsumerNumber,
			c.cfg.MaxWorkingProducers))
	}

	filename.WriteString(c.ext)

	filePath := filepath.Join(c.cfg.FilesDirName, filename.String())

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
