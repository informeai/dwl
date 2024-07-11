package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

// Download is struct for execute download
type Download struct {
	url     string
	outFile string
}

// NewDownload return instance of download
func NewDownload() *Download {
	return &Download{}
}

// Start execute request for download
func (d *Download) Start(url, outFile string) error {
	response, err := d.fetch(url)
	if err != nil {
		return err
	}
	file, err := d.prepareFile(outFile)
	if err != nil {
		return err
	}
	if err := d.prepareProgressBar(response, file); err != nil {
		return err
	}
	d.finished()
	return nil
}

// finished print concluded download
func (d *Download) finished() {
	fmt.Println("Download Concluded")
}

// fetch execute request with url
func (d *Download) fetch(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// prepareFile prepare file for write
func (d *Download) prepareFile(outFile string) (*os.File, error) {
	f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0o664)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// prepareProgressBar execute configuration the progress bar
func (d *Download) prepareProgressBar(res *http.Response, file *os.File) error {
	defer res.Body.Close()
	defer file.Close()
	bar := progressbar.DefaultBytes(
		res.ContentLength,
		"downloading",
	)
	if _, err := io.Copy(io.MultiWriter(file, bar), res.Body); err != nil {
		return err
	}
	return nil
}
