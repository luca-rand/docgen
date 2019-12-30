package plugins

import "html/template"

// Logger handles logging
type Logger interface {
	Debug(...string)
	Info(m ...string)
	Fatal(...string)
}

// ProjectConfig is the global project config for docgen
type ProjectConfig struct {
	OutputFolder string   `yaml:"output_folder"`
	Exclude      []string `yaml:"exclude"`
}

// PluginConfig is the docgen config for a specific plugin
type PluginConfig map[string]interface{}

// Meta is the information about a project + config that is passed to the hooks
type Meta struct {
	Logger  Logger
	Project ProjectConfig
	Plugin  PluginConfig
}

// Page is a page that will be rendered
type Page interface {
	Title() string
	Slug() string

	Data() map[string]interface{}
	Content() string
}

// Plugin is a docgen thing that extends functionality
type Plugin struct {
	InitHook func(meta *Meta)

	TitleHook      func(meta *Meta, page Page) (title string)
	TemplatingHook func(meta *Meta, page Page) (data interface{}, funcs template.FuncMap)
	RenderHook     func(meta *Meta, page Page) (head string, before string, after string)
}
