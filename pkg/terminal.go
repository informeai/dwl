package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

// Term is struct for actions in terminal
type Term struct {
	version string
	app     *cli.App
}

// NewTerm return instance of term
func NewTerm() *Term {
	return &Term{
		version: "1.0.0",
	}
}

// isClosed verify signals
func (t *Term) isClosed() {
	interruptChan := make(chan os.Signal, 1)

	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		<-interruptChan
		signal.Stop(interruptChan)
		os.Exit(1)
	}()
}

// Prepare configure term for action
func (t *Term) Prepare(downloader IDownload) error {
	t.isClosed()
	app := &cli.App{
		Name:  "Dwl",
		Usage: "Downloader for files.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output for file downloaded",
				Aliases: []string{"o"},
			},
			&cli.StringFlag{
				Name:    "url",
				Usage:   "Url for download",
				Aliases: []string{"u"},
			},
		},
		Version: t.version,
		Authors: []*cli.Author{
			{
				Name:  "Wellington Gadelha",
				Email: "contato.informeai@gmail.com",
			},
		},
		Action: func(c *cli.Context) error {
			return downloader.Start(c.String("url"), c.String("output"))
		},
	}
	t.app = app
	return nil
}

// Run execute actions
func (t *Term) Run() error {
	if err := t.app.Run(os.Args); err != nil {
		return err
	}
	return nil
}
