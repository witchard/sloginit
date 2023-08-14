// auto initialises the default log/slog using sloginit defaults.
package auto

import (
	"log/slog"

	"github.com/witchard/sloginit"
)

func init() {
	slog.SetDefault(sloginit.Logger())
}
