package main

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/witchard/sloginit"
)

func main() {
	tty := isatty.IsTerminal(os.Stderr.Fd())

	log := sloginit.Logger()
	log.Debug("debug", "tty", tty)
	log.Info("info", "tty", tty)
	log.Warn("warn", "tty", tty)
	log.Error("error", "tty", tty)
}
