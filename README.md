# dwl

Downloader files written in go

### Build

Clone this repository and build with

```
go build -o dwl cmd/main.go
```

### Usage

Single example of usage

```
dwl -u http://example.com/file.txt -o file.txt
```

#### Help

Help the commands

```
dwl help
```

### Libs

[cli](https://github.com/urfave/cli) - actions e args parse.

[progressbar](https://github.com/schollz/progressbar) - the progress bar in terminal

by informeai :heart:
