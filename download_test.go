package canopy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewDownload(t *testing.T) {
	d := NewDownload("url", "filepath")
	if d.filepath != "filepath" && d.url != "url" {
		t.Errorf("NewDownload argument mismatch: %v", d)
	}
}

func TestDownloadGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			http.NotFound(w, r)
		}
		fmt.Fprintln(w, "this is your test file")
	}))
	defer ts.Close()

	// a good download
	d := NewDownload(ts.URL, "test_file")
	if n, err := d.Get(); n == 0 || err != nil {
		t.Errorf("Download.Get (good) unexpected returns: %v, %v", n, err)
	}

	// a bad filename
	d = NewDownload(ts.URL, "a/bad/test_file")
	if n, err := d.Get(); n != 0 || err == nil {
		t.Errorf("Download.Get (bad filename) unexpected returns: %v, %v", n, err)
	}

	// a not found url
	d = NewDownload(ts.URL+"/not-found-url", "test_file")
	if n, err := d.Get(); n != 0 || err == nil {
		t.Errorf("Download.Get (not found url) unexpected returns: %v, %v", n, err)
	}
}
