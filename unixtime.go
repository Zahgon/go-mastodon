package mastodon

import (
	"time"
)

type Unixtime time.Time

func (t *Unixtime) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
