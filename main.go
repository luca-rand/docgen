package main

import "fmt"

import "path/filepath"

import "strings"

func main() {
	logger := &Logger{level: debug}
	logger.Debug("created logger")

	config := openConfig(logger)
	logger.Debug("loaded config:", fmt.Sprint(config))

	loadedPlugins := openAndInitPlugins(config, logger)
	logger.Info(fmt.Sprintf("Loaded %d plugins", len(loadedPlugins)))

	sources := getSources(config, logger)
	logger.Info(fmt.Sprintf("Found %d sources", len(sources)))
	for _, source := range sources {
		slug := filepath.Dir(source)
		base := filepath.Base(source)
		base = strings.TrimSuffix(base, filepath.Ext(base))
		if base != "index" {
			slug = filepath.Join(slug, base)
		}
		if slug == "README" {
			slug = ""
		}
		slug = filepath.Clean(slug)

		renderPage(logger, config, loadedPlugins, &Page{
			logger:     logger.PagePrefix(slug),
			title:      slug,
			slug:       slug,
			sourcePath: source,
		})

	}

}
