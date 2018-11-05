package proxy

import (
	"time"

	"github.com/hellofresh/janus/pkg/router"
)

// RegisterOption represents the register options
type RegisterOption func(*Register)

// WithRouter sets the router
func WithRouter(router router.Router) RegisterOption {
	return func(r *Register) {
		r.router = router
	}
}

// WithFlushInterval sets the Flush interval for copying upgraded connections
func WithFlushInterval(d time.Duration) RegisterOption {
	return func(r *Register) {
		r.flushInterval = d
	}
}

// WithIdleConnectionsPerHost sets idle connections per host option
func WithIdleConnectionsPerHost(value int) RegisterOption {
	return func(r *Register) {
		r.idleConnectionsPerHost = value
	}
}

// WithIdleConnTimeout sets the maximum amount of time an idle
// (keep-alive) connection will remain idle before closing
// itself.
func WithIdleConnTimeout(d time.Duration) RegisterOption {
	return func(r *Register) {
		r.idleConnTimeout = d
	}
}
