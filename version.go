package registry

const apiVersionHeader = "Docker-Distribution-Api-Version"

// APIVersion performs a version check on the registry API.
func (r *Registry) APIVersion() (string, error) {
	url := r.host + "/v2/"
	resp, err := r.client.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Header.Get(apiVersionHeader), nil
}
