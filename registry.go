package registry

import (
	"net/http"
	"time"
)

// Registry represents a Docker Registry.
type Registry struct {
	host   string
	client *http.Client
}

// NewRegistry returns a new instance of Registry with connection properties to host.
func NewRegistry(host string, timeout time.Duration) *Registry {
	return &Registry{
		host: host,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}
