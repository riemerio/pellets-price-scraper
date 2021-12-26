package prices

import "time"

type Price struct {
	RecordId  int
	Price     float32
	Timestamp time.Time
}
