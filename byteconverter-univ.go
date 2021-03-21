// Copyright 2021 Mia Metzler. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//The package go-byte-converter provides dead simple interfaces to convert byte array representation of data back into its original datatype.

package byteconverter

//LittleEndian binary converter
import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	BigEndian    = iota
	LittleEndian = iota
)

//Converts 4 byte array representation of a float32 back into a float32.
func Read_float32(endian int, bytes []byte) float32 {
	var bits = uint32(0)
	if endian == BigEndian {
		bits = binary.BigEndian.Uint32(bytes)
	} else {
		bits = binary.LittleEndian.Uint32(bytes)
	}
	float := math.Float32frombits(bits)
	return float
}

//Converts 8 byte array representation of a float64 back into a float64.
func Read_float64(endian int, bytes []byte) float64 {
	var bits = uint64(0)
	if endian == BigEndian {
		bits = binary.BigEndian.Uint64(bytes)
	} else {
		bits = binary.LittleEndian.Uint64(bytes)
	}

	float := math.Float64frombits(bits)
	return float
}

//Converts 1 byte array representation of a int8 back into a int8.
func Read_int8(endian int, data []byte) (ret int8) {
	buf := bytes.NewBuffer(data)
	if endian == BigEndian {
		binary.Read(buf, binary.BigEndian, &ret)
	} else {
		binary.Read(buf, binary.LittleEndian, &ret)
	}

	return
}

//Converts 2 byte array representation of a int16 back into a int16.
func Read_int16(endian int, data []byte) (ret int16) {
	buf := bytes.NewBuffer(data)
	if endian == BigEndian {
		binary.Read(buf, binary.BigEndian, &ret)
	} else {
		binary.Read(buf, binary.LittleEndian, &ret)
	}

	return
}

//Converts 4 byte array representation of a int32 back into a int32.
func Read_int32(endian int, data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	if endian == BigEndian {
		binary.Read(buf, binary.BigEndian, &ret)
	} else {
		binary.Read(buf, binary.LittleEndian, &ret)
	}
	return
}

//Converts 8 byte array representation of a int64 back into a int64.
func Read_int64(endian int, data []byte) (ret int64) {
	buf := bytes.NewBuffer(data)

	if endian == BigEndian {
		binary.Read(buf, binary.BigEndian, &ret)
	} else {
		binary.Read(buf, binary.LittleEndian, &ret)
	}
	return
}

//Converts 1 byte array representation of a uint8 back into a uint8.
func Read_uint8(endian int, data []byte) uint8 {
	return data[0] //byte is an alias for uint8 and is equivalent to uint8 in all ways.
}

//Converts 2 byte array representation of a uint16 back into a uint16.
func Read_uint16(endian int, data []byte) uint16 {
	if endian == BigEndian {
		return binary.BigEndian.Uint16(data)
	} else {
		return binary.LittleEndian.Uint16(data)
	}

}

//Converts 4 byte array representation of a uint32 back into a uint32.
func Read_uint32(endian int, data []byte) uint32 {
	if endian == BigEndian {
		return binary.BigEndian.Uint32(data)
	} else {
		return binary.LittleEndian.Uint32(data)
	}
}

//Converts 8 byte array representation of a uint64 back into a uint64.
func Read_uint64(endian int, data []byte) uint64 {
	if endian == BigEndian {
		return binary.BigEndian.Uint64(data)
	} else {
		return binary.LittleEndian.Uint64(data)
	}
}
