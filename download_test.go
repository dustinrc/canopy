package canopy

import "testing"

func TestDownload(t *testing.T) {
	d := NewDownload("url", "filepath")
	if d.filepath != "filepath" && d.url != "url" {
		t.Errorf("NewDownload argument mismatch: %v", d)
	}
}
