# PDFTOJPG

No frils CLI command to convert pdf to jpg image(s). Extended the script from: https://golangcode.com/convert-pdf-to-jpg/

## Prerequisite
- GO 1.17
- imagemagick

## Build
```shell
# install imagemagick
brew install ghostscript imagemagick
# allow CFLAG to compile main.go
export CGO_CFLAGS_ALLOW='-Xpreprocessor'

go build # will output `pdftojpg` binary
```

## Usage

```shell
# pass in pdf source, jpg output, number of page
./pdftojpg 2023.3.10.pdf 2023.3.10 440
```
