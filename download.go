package canopy

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download associates a url and the filepath it will be stored.
type Download struct {
	filepath string
	url      string
}

// NewDownload creates a new Download instance.
func NewDownload(url, filepath string) *Download { return &Download{filepath, url} }

// Get retrieves the content from url and stores it at filepath.
// It returns the number of bytes copied and the first error if any.
func (d *Download) Get() (n int64, err error) {
	file, err := os.Create(d.filepath)
	if err != nil {
		return
	}
	defer file.Close()

	resp, err := http.Get(d.url)
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("non-OK status code: %s", resp.Status)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()

	n, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	return
}
