package availabilitychecks

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RegistryTarget map[string]string

type RegistryAvailabilityChecker struct {
	Targets []RegistryTarget
}

type RegistryConfig struct {
	Targets []RegistryTarget
}

// NewRegistryAvailabilityChecks accepts a Config and attempts to ascert its type
// as RegistryConfig. If ascertation is true a NewRegistryAvailabilityChecker is returned.
// Else, an error is returned advising failed ascertation.
func NewRegistryAvailabilityChecker(c Config) (*RegistryAvailabilityChecker, error) {
	// its ok to return 0 length data
	if data, ok := c.(RegistryConfig); ok {
		return &RegistryAvailabilityChecker{Targets: data.Targets}, nil
	}
	return &RegistryAvailabilityChecker{}, fmt.Errorf("Attempt to get Registry implementation failed ascertation as RegistryConfig")
}

// HTTPAvailabilityCheck accepts a slice of HTTP targets accompanied by a timeout
// and asynchronously checks the targets to deem them available. If any of the targets
// are deemed unhealthy, all running routines are cancelled and an error is returned.
func (r RegistryAvailabilityChecker) AvailabilityCheck(timeout time.Duration) error {
	client := http.Client{
		Timeout: (timeout * time.Second),
	}

	var RegistryAvailabilityError error

	// Utilize a WaitGroup to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// Set the WaitGroup counter to the length of targets. The WaitGroup
	// can then block completion of the call until all goroutines have completed
	// by calling wg.Done
	wg.Add(len(r.Targets))

	// Add context here? that spans each goroutine and cancels if failures?

	for _, url := range r.Targets {

		go func(url string) {
			err := retry(3, time.Second, func() error {
				// registry logic here
			})

			// protect with mutex
			// if this is ever not nil, cancel all routines with mutex?
			if err != nil {
				RegistryAvailabilityError = err
			}

			// drops counter by 1
			wg.Done()

		}(url)
	}
	// Waits for WaitGroup counter to be 0.
	wg.Wait()
	return RegistryAvailabilityError
}
