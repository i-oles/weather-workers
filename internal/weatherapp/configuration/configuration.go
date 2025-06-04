package configuration

import "encoding/json"

type Configuration struct {
	SourceFileName           string
	APIURL                   string
	LogProducedMsg           bool
	LogConsumedMsg           bool
	LogResults               bool
	AnalysisDurationInMonths int
	Mode                     string
	MockAPI                  bool
	ExecutionRepeatCount     int
	PerformanceTest          bool
	FilesDirName             string
	ConsumerNumber           int
	ProducerNumber           int
	MaxWorkingProducers      int
}

func (c *Configuration) Pretty() string {
	cfgPretty, _ := json.MarshalIndent(c, "", "  ")

	return string(cfgPretty)
}
