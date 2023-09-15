# sloginit

An opinionated initialisation for the slog package.

Currently works as follows:
* If output is a TTY will use github.com/lmittmann/tint for logging
* If output is not it will use JSON logging

Settings can be overridden by environment variables:
* `LOG_LEVEL` - the level to log at
* `LOG_JSON` - force JSON logger if set to 1/ON/YES/TRUE, otherwise use console logger

Supports auto initialisation of the slog default logger:

```go
package main

import (
	"log/slog"

	_ "github.com/witchard/sloginit/auto"
)

func main() {
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}
```

Or can be initialised directly so you can provide setup options:

```go
package main

import (
	"log/slog"

	"github.com/witchard/sloginit"
)

func main() {
	sloginit.SetDefault(sloginit.JSON)
	slog.Info("hi", "format", "json")
}
```
