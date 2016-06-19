package registry

import (
	"io"
	"net/http"
)

// Catalog lists all the repositories in the registry.
func (r *Registry) Catalog() (io.ReadCloser, error) {
	url := r.Host + "/v2/_catalog"
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
