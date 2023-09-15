// auto initialises the default log/slog using sloginit defaults.
package auto

import (
	"github.com/witchard/sloginit"
)

func init() {
	sloginit.SetDefault()
}
