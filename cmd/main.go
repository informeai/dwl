package main

import (
	"log"

	"github.com/informeai/dwl/pkg"
)

func main() {
	downloader := pkg.NewDownload()
	term := pkg.NewTerm()
	if err := term.Prepare(downloader); err != nil {
		log.Fatal(err)
	}
	if err := term.Run(); err != nil {
		log.Fatal(err)
	}
}
