package config

import (
	"flag"
	"log/slog"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

const (
	defaultConfigName = "dev"
	ext               = ".json"
)

func GetConfig(cfgBasePath string, cfg any) error {
	profile := flag.String("profile", "", "config name to load from")
	flag.Parse()

	configName := defaultConfigName
	if *profile != "" {
		configName = *profile
	}

	cfgPath := filepath.Join(cfgBasePath, configName+ext)

	err := gonfig.GetConf(cfgPath, cfg)
	if err != nil {
		slog.Error("could not get config file", "error", err)
	}

	return nil
}
