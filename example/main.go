package main

import (
	"os"
	"log/slog"

	"github.com/mattn/go-isatty"
	_ "github.com/witchard/sloginit/auto"
)

func main() {
	tty := isatty.IsTerminal(os.Stderr.Fd())

	slog.Debug("debug", "tty", tty)
	slog.Info("info", "tty", tty)
	slog.Warn("warn", "tty", tty)
	slog.Error("error", "tty", tty)
}
