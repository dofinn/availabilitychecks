package availabilitychecks

import (
	"time"
)

type Config interface{}

type AvailabilityChecker interface {
	AvailabilityCheck(timeout time.Duration) error
}
