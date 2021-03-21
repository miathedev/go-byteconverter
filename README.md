# go-byte-converter

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

## Why did you create this library/module?

Because im lazy. Ive got some TheThingsNetwork IoT nodes that send me raw mimimum data (to save bandwith over LoRa). I have to convert all data back myself (because the payload format feature from ttn is not powerful enough). I had to do this in many project. So this library has been born.

## Usage

Please take a look at the _test.go file


