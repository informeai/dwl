package tests

import (
	"log"
	"testing"

	"github.com/informeai/dwl/pkg"
)

func TestNewTerm(t *testing.T) {
	term := pkg.NewTerm()
	if term == nil {
		t.Errorf("TestNewTerm: expect(!nil) got(%v)\n", term)
	}
}

func TestTermPrepare(t *testing.T) {
	term := pkg.NewTerm()
	if term == nil {
		t.Errorf("TestTermPrepare: expect(!nil) got(%v)\n", term)
	}
	downloader := pkg.NewDownload()
	if err := term.Prepare(downloader); err != nil {
		t.Errorf("TestTermPrepare: expect(nil) got(%s)\n", err.Error())
	}
	log.Printf("downloader: %+v\n", downloader)
}

func TestTermRun(t *testing.T) {
	term := pkg.NewTerm()
	if term == nil {
		t.Errorf("TestTermRun: expect(!nil) got(%v)\n", term)
	}
	downloader := pkg.NewDownload()
	if err := term.Prepare(downloader); err != nil {
		t.Errorf("TestTermRun: expect(nil) got(%s)\n", err.Error())
	}
  if err := term.Run(); err != nil{
		t.Errorf("TestTermRun: expect(nil) got(%s)\n", err.Error())
  }
}
