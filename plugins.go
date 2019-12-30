package main

import "github.com/luca-rand/docgen/plugins"

// availablePlugins is the list of loaded plugins
var availablePlugins = make(map[string]plugins.Plugin)

func init() {
	availablePlugins["gitlab"] = plugins.Plugin{}
}

type pluginWithMeta struct {
	plugins.Plugin
	meta *plugins.Meta
}

// openAndInitPlugins opens and initalizes the plugins
func openAndInitPlugins(config *Config, logger *Logger) map[string]pluginWithMeta {
	usedPlugins := make(map[string]pluginWithMeta)

	for pluginName, pluginConfig := range config.Plugins {
		var ok bool
		var plugin plugins.Plugin
		if plugin, ok = availablePlugins[pluginName]; !ok {
			logger.Fatal("plugin", pluginName, "could not be loaded as it is not available")
		}
		usedPlugins[pluginName] = pluginWithMeta{
			Plugin: plugin,
			meta: &plugins.Meta{
				Logger:  logger.PluginPrefix(pluginName),
				Project: config.ProjectConfig,
				Plugin:  pluginConfig,
			},
		}
		if usedPlugins[pluginName].InitHook != nil {
			usedPlugins[pluginName].InitHook(usedPlugins[pluginName].meta)
		}
	}

	return usedPlugins
}
