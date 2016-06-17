package registry

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockAPI))
	defer server.Close()

	fixture := &Registry{Host: server.URL}
	expected := "registry/2.0"
	actual, err := fixture.APIVersion()
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Errorf("Expected registry API version to be %q, but got %q", expected, actual)
	}
}

func mockAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Docker-Distribution-Api-Version", "registry/2.0")
	w.Header().Set("Content-Length", "2")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	_, err := fmt.Fprint(w, "{}")
	if err != nil {
		log.Fatal(err)
	}
}
