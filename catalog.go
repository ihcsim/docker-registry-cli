package registry

import "io"

// Catalog lists all the repositories in the registry.
func (r *Registry) Catalog() (io.ReadCloser, error) {
	url := r.host + "/v2/_catalog"
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
