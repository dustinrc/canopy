package canopy

import (
	"io"
	"net/http"
	"os"
)

type Download struct {
	filepath string
	url      string
}

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
