package canopy

import (
	"fmt"
	. "launchpad.net/gocheck"
	"net/http"
	"net/http/httptest"
)

func (s *S) TestNewDownload(c *C) {
	d := NewDownload("url", "filepath")
	c.Assert(d.filepath, Equals, "filepath")
	c.Assert(d.url, Equals, "url")
}

func (s *S) TestDownloadGet(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			http.NotFound(w, r)
		}
		fmt.Fprintln(w, "this is your test file")
	}))
	defer ts.Close()

	// a good download
	d := NewDownload(ts.URL, "test_file")
	n, err := d.Get()
	c.Assert(n, Not(Equals), 0)
	c.Assert(err, IsNil)

	// a bad filename
	d = NewDownload(ts.URL, "a/bad/test_file")
	n, err = d.Get()
	c.Assert(n, Equals, int64(0))
	c.Assert(err, NotNil)

	// a not found url
	d = NewDownload(ts.URL+"/not-found-url", "test_file")
	n, err = d.Get()
	c.Assert(n, Equals, int64(0))
	c.Assert(err, NotNil)
}
