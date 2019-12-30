package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/markbates/pkger"
)

var t *template.Template

func openTemplate(logger *Logger) {
	logger.Debug("opening base template")
	file, err := pkger.Open("/static/page.html")
	if err != nil {
		logger.Fatal("failed to open base template:", err.Error())
	}
	defer file.Close()

	logger.Debug("reading base template")
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Fatal("failed to read base template:", err.Error())
	}

	logger.Debug("parsing base template")
	t, err = template.New("").Parse(string(fileData))
	if err != nil {
		logger.Fatal("failed to parse base template:", err.Error())
	}
}

// renderPage renders a single page
func renderPage(logger *Logger, config *Config, loadedPlugins map[string]pluginWithMeta, page *Page) {
	for name, plugin := range loadedPlugins {
		if plugin.TitleHook != nil {
			page.logger = logger.PluginPrefix(name)
			page.title = plugin.TitleHook(plugin.meta, page)
		}
	}
	page.logger = logger

	data := make(map[string]interface{})
	data["Page"] = page
	funcs := make(map[string]interface{})

	for name, plugin := range loadedPlugins {
		page.logger = logger.PluginPrefix(name)
		if plugin.TemplatingHook != nil {
			pluginData, pluginFuncs := plugin.TemplatingHook(plugin.meta, page)
			data[name] = pluginData
			for k, v := range pluginFuncs {
				funcs[name+"_"+k] = v
			}
		}
	}
	page.logger = logger

	var head = ""
	var before = ""
	var after = ""

	for name, plugin := range loadedPlugins {
		page.logger = logger.PluginPrefix(name)
		if plugin.RenderHook != nil {
			pluginHead, pluginBefore, pluginAfter := plugin.RenderHook(plugin.meta, page)
			head += pluginHead
			before += pluginBefore
			after += pluginAfter
		}
	}
	page.logger = logger

	if t == nil {
		logger.Debug("base template not initalized")
		openTemplate(logger)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	t, err := t.Clone()
	if err != nil {
		logger.Fatal("failed to clone base template:", err.Error())
	}

	_, err = t.New("__head__").Parse(head)
	if err != nil {
		logger.Fatal("failed to parse __head__ template:", err.Error())
	}
	_, err = t.New("__beforeContent__").Parse(before)
	if err != nil {
		logger.Fatal("failed to parse __beforeContent__ template:", err.Error())
	}
	_, err = t.New("__content__").Parse(string(markdown.ToHTML([]byte(page.Content()), parser, renderer)))
	if err != nil {
		logger.Fatal("failed to parse __content__ template:", err.Error())
	}
	_, err = t.New("__afterContent__").Parse(after)
	if err != nil {
		logger.Fatal("failed to parse __afterContent__ template:", err.Error())
	}

	path := filepath.Join(config.OutputFolder, page.Slug(), "index.html")
	err = os.MkdirAll(filepath.Dir(path), os.ModeDir)
	if err != nil {
		logger.Fatal("could not create directory structure for output file:", err.Error())
	}

	out, err := os.Create(path)
	if err != nil {
		logger.Fatal("could not open the output file:", err.Error())
	}
	defer out.Close()

	err = t.Funcs(funcs).Execute(out, data)
	if err != nil {
		logger.Fatal("failed to execute template:", err.Error())
	}
}
