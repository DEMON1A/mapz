# Mapz
Mapz is a tool written in Go to validate the existence of JavaScript map files on websites.

## Installation
To install Mapz, you can use the `go install` command with the following:
```bash
go install -v github.com/DEMON1A/mapz/cmd/mapz@latest
```
Make sure you have Go installed and your `$GOPATH/bin` directory is added to your `PATH` environment variable.

## Usage
```css
Usage of mapz.exe:
  -fast
        Perform a fast scan by disabling a validate rule, but it may result in some false positives
  -file string
        Path to your file that contains JavaScript URLs
  -url string
        The URL you want to validate the map existance for
  -verbose
        Verbose mode to show additional information regarding the output
  -workers int
        Number of workers to use for concurency (default 4)
```

## Disclaimer
**Disclaimer:** The Mapz tool is intended for educational and legitimate testing purposes only. Be aware that scanning websites or accessing resources without proper authorization may violate terms of service and local laws. The author of this tool is not responsible for any misuse or illegal actions conducted with this tool. Use Mapz responsibly and at your own risk.
