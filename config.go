package main

import "github.com/luca-rand/docgen/plugins"

import "os"

import "gopkg.in/yaml.v2"

// Config object used to configure the generator
type Config struct {
	plugins.ProjectConfig `yaml:",inline"`
	Plugins               map[string]plugins.PluginConfig `yaml:"plugins"`
}

func openConfig(logger *Logger) *Config {
	logger.Debug("opening docgen.yml")
	file, err := os.Open("docgen.yml")
	if err != nil {
		logger.Fatal("failed to open docgen.yml:", err.Error())
	}
	defer file.Close()

	config := &Config{}
	logger.Debug("parsing docgen.yml")
	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		logger.Fatal("failed to parse docgen.yml:", err.Error())
	}

	if config.OutputFolder == "" {
		logger.Fatal("An output_folder has to be specified in docgen.yml.")
	}

	return config
}
