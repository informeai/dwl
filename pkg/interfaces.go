package pkg

// IDownload is interface for downloader
type IDownload interface {
	Start(url, outFile string) error
}
