package tests

import (
	"log"
	"testing"

	"github.com/informeai/dwl/pkg"
)

func TestNewDownload(t *testing.T) {
	download := pkg.NewDownload()
	if download == nil {
		t.Errorf("TestNewDownload: expect(!nil) got(%v)\n", download)
	}
	log.Printf("download: %+v\n", download)
}

func TestDownloadStart(t *testing.T) {
	download := pkg.NewDownload()
	if download == nil {
		t.Errorf("TestNewDownload: expect(!nil) got(%v)\n", download)
	}
	if err := download.Start("https://example.com/file.txt", "file.txt"); err != nil {
		t.Errorf("TestNewDownload: expect(nil) got(%s)\n", err.Error())
	}
}
