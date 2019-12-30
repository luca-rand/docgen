package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/lucacasonato/frontmatter"
	"gopkg.in/yaml.v2"
)

// Page is the implementation of plugins.Page
type Page struct {
	logger     *Logger
	sourcePath string
	title      string
	slug       string
}

func (p *Page) log() *Logger {
	return p.logger.PagePrefix(p.slug)
}

// Title is the name of the page
func (p *Page) Title() string {
	return p.title
}

// Slug is the output url of this page
func (p *Page) Slug() string {
	return p.slug
}

func (p *Page) open() io.ReadCloser {
	reader, err := os.Open(p.sourcePath)
	if err != nil {
		p.log().Fatal(err.Error())
	}

	return reader
}

func (p *Page) openAndSplit() (string, string) {
	file := p.open()
	defer file.Close()
	fm, content, err := frontmatter.Parse(file, "---")
	if err != nil {
		if err == frontmatter.ErrMissingDelimiter {
			file := p.open()
			defer file.Close()
			var data []byte
			data, err = ioutil.ReadAll(file)
			if err == nil {
				return "", string(data)
			}
		}
		p.log().Fatal(err.Error())
	}

	return fm, content
}

// Data is the parsed content of the yaml frontmatter if it exists
func (p *Page) Data() map[string]interface{} {
	frontmatter, _ := p.openAndSplit()
	data := map[string]interface{}{}

	err := yaml.Unmarshal([]byte(frontmatter), &data)
	if err != nil {
		p.log().Fatal(err.Error())
	}

	return data
}

// Content is the page content without the yaml frontmatter if it exists
func (p *Page) Content() string {
	_, data := p.openAndSplit()
	return data
}
