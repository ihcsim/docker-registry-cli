package registry

import (
	"net/http"
	"time"
)

const (
	timeout          = time.Second * 2
	apiVersionHeader = "Docker-Distribution-Api-Version"
)

// APIVersion performs a version check on the registry API.
func (r *Registry) APIVersion() (string, error) {
	url := r.Host + "/v2/"
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Header.Get(apiVersionHeader), nil
}
