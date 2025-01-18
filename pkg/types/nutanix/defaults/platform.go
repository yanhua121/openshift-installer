package defaults

import (
	"time"

	"github.com/openshift/installer/pkg/types/nutanix"
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *nutanix.Platform) {
	// If PrismAPICallTimeout is not configured, the default value of 5 minutes will be used
	// as the prism-api call timeout.
	timeout := nutanix.DefaultPrismAPICallTimeout
	if p.PrismAPICallTimeout == nil {
		p.PrismAPICallTimeout = &timeout
	} else {
		timeout = *p.PrismAPICallTimeout
	}

	nutanix.PrismAPICallTimeoutDuration = time.Duration(timeout) * time.Minute
}
