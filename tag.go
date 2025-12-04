package git

import (
	"time"
)

// Tag houses some common info about tags.
type Tag struct {
	Name string
	Hash Hash
	Date time.Time
}
