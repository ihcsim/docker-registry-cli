package registry

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const expected = `{"repositories":["golang:1.5.4", "golang:1.6.1", "golang:1.6.2"]}`

func TestCatalog(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(mockCatalog))
	defer s.Close()

	r := &Registry{Host: s.URL}
	res, err := r.Catalog()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Close()

	var b []byte
	buf := bytes.NewBuffer(b)
	n, err := buf.ReadFrom(res)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if n == 0 {
		t.Errorf("Expected non-zero bytes content")
	}

	if buf.String() != expected {
		t.Errorf("Expected repository list to be %q, but got %q", expected, buf.String())
	}
}

func mockCatalog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(expected))
}
