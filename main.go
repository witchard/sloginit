// sloginit provides simple initialisation for the slog package.
package sloginit

import (
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

type opts struct {
	console bool
	level   slog.Level
	prefix  string
}

// SlogOpt provides functional options to configure the logger.
type SlogOpt func(*opts)

// Console is a SlogOpt that forces output to the terminal.
func Console(o *opts) {
	o.console = true
}

// JSON is a SlogOpt that forces output as JSON.
func JSON(o *opts) {
	o.console = false
}

// Level is a SlogOpt to set the log level.
func Level(l slog.Level) SlogOpt {
	return func(o *opts) {
		o.level = l
	}
}

// EnvPrefix is a SlogOpt that sets the env prefix for reading environment options.
func EnvPrefix(prefix string) SlogOpt {
	return func(o *opts) {
		o.prefix = prefix
	}
}

func setOpts(options ...SlogOpt) opts {
	// First we setup option defaults
	//  - Console if the output is a terminal (otherwise JSON)
	//  - Level is INFO
	o := opts{
		console: isatty.IsTerminal(os.Stderr.Fd()),
		level:   slog.LevelInfo,
		prefix:  "LOG",
	}

	// Run options so we know the environment prefix
	for _, opt := range options {
		opt(&o)
	}

	// Extract environment settings
	json := os.Getenv(o.prefix + "_JSON")
	if json != "" {
		o.console = !slices.Contains([]string{"TRUE", "1", "YES", "ON"}, strings.ToUpper(json))
	}
	level := os.Getenv(o.prefix + "_LEVEL")
	if level != "" {
		switch strings.ToUpper(level) {
		case "D", "DEBUG":
			o.level = slog.LevelDebug
		case "I", "INFO", "INFORMATION":
			o.level = slog.LevelInfo
		case "W", "WARN", "WARNING":
			o.level = slog.LevelWarn
		case "E", "ERR", "ERROR":
			o.level = slog.LevelError
		}
	}

	// Then we run options again to override environment
	for _, opt := range options {
		opt(&o)
	}
	return o
}

func Logger(options ...SlogOpt) *slog.Logger {
	o := setOpts(options...)
	var logger *slog.Logger
	if o.console {
		logger = slog.New(
			tint.NewHandler(os.Stderr, &tint.Options{
				Level:      o.level,
				TimeFormat: time.Kitchen,
			}),
		)
	} else {
		logger = slog.New(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
				Level: o.level,
			}),
		)
	}
	return logger
}
