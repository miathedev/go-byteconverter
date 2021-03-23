# go-byte-converter

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/miathedev/go-byteconverter)

Dead simple conversion between byte arrays and all common data types. Useful for raw IoT payloads.

## Supported Datatypes

### uint

- uint64 (untested)
- uint32
- uint16
- uint8

### int

- int64 (untested)
- int32
- int16
- int8

### float

- float64 (untested)
- float32

### string

- byte of ascii elements

## Why did you create this library/module?

Because im lazy. Ive got some TheThingsNetwork IoT nodes that send me raw mimimum data (to save bandwith over LoRa). I have to convert all data back myself (because the payload format feature from ttn is not powerful enough). I had to do this in many projects. Thats why this library has been born.

## Usage

Please take a look at the _test.go file


