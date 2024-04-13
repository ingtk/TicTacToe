package pkg

import "time"

type Clock interface {
	Now() time.Time
}
