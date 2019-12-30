package main

import "fmt"

import "os"

import "strings"

type logLevel int

// The log level for the logger
const (
	debug logLevel = iota
	info
	fatal
)

// Logger logs things to the console
type Logger struct {
	level  logLevel
	prefix string
}

// Prefix creates a new logger with that prefix
func (l *Logger) Prefix(prefix string) *Logger {
	if l.prefix != "" {
		prefix = l.prefix + " " + prefix
	}

	return &Logger{
		prefix: prefix,
	}
}

// PluginPrefix creates a new logger with a prefix for plugins
func (l *Logger) PluginPrefix(plugin string) *Logger {
	return l.Prefix("[plugin " + plugin + "]")
}

// PagePrefix creates a new logger with a prefix for pages
func (l *Logger) PagePrefix(page string) *Logger {
	return l.Prefix("[page " + page + "]")
}

// Debug prints a debug message to the log
func (l *Logger) Debug(messages ...string) {
	if l.level <= debug {
		fmt.Println("[debug]", l.prefix, strings.Join(messages, " "))
	}
}

// Info prints a informational message to the log
func (l *Logger) Info(messages ...string) {
	if l.level <= info {
		fmt.Println("[info]", l.prefix, strings.Join(messages, " "))
	}
}

// Fatal prints a fatal error to the log and then exits
func (l *Logger) Fatal(messages ...string) {
	if l.level <= fatal {
		fmt.Println("[fatal]", l.prefix, strings.Join(messages, " "))
		os.Exit(1)
	}
}
