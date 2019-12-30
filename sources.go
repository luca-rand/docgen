package main

import "github.com/bmatcuk/doublestar"

func getSources(config *Config, logger *Logger) []string {
	sources, err := doublestar.Glob("docs/**/*.md")
	if err != nil {
		logger.Fatal("failed to get all potential markdown files:", err.Error())
	}

	for _, exclude := range config.Exclude {
		var tempSources []string
		for _, source := range sources {
			match, err := doublestar.Match(exclude, source)
			if err != nil {
				logger.Fatal("exclude matching failed:", err.Error())
			}
			if !match {
				tempSources = append(tempSources, source)
			}
		}
		sources = tempSources
	}

	sources = append([]string{"README.md"}, sources...)

	return sources
}
