package views

import (
	"fmt"
	"time"
)

// AssetVersion is a timestamp generated at app start for cache busting.
// Used in templates to append ?v=<timestamp> to URLs.
var AssetVerion = fmt.Sprintf("%d", time.Now().Unix())
