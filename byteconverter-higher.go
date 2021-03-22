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

//Converts byte array representation back to its matching type. The type is figured out by the size of the array.
//Eg. size 1: int8, size2: int16, size4: int32, ...
//Returns 0 if array dimensions didnt match
func Read_int(endian int, bytes []byte) (ret int64) {
	var size = len(bytes)
	switch size {
	case 1:
		return int64(Read_int8(endian, bytes))
	case 2:
		return int64(Read_int16(endian, bytes))
	case 4:
		return int64(Read_int32(endian, bytes))
	case 8:
		return Read_int64(endian, bytes)
	default:
		return 0
	}
}

//Converts byte array representation back to its matching type. The type is figured out by the size of the array.
//Eg. size 1: uint8, size2: uint16, size4: uint32, ...
//Returns 0 if array dimensions didnt match
func Read_uint(endian int, bytes []byte) (ret uint64) {
	var size = len(bytes)
	switch size {
	case 1:
		return uint64(Read_uint8(endian, bytes))
	case 2:
		return uint64(Read_uint16(endian, bytes))
	case 4:
		return uint64(Read_uint32(endian, bytes))
	case 8:
		return Read_uint64(endian, bytes)
	default:
		return 0
	}
}

//Converts byte array representation back to its matching type. The type is figured out by the size of the array.
//Eg. size 4: float32, size 8: float64
//Returns 0 if array dimensions didnt match
func Read_float(endian int, bytes []byte) (ret float64) {
	var size = len(bytes)
	switch size {
	case 4:
		return float64(Read_float32(endian, bytes))
	case 8:
		return Read_float64(endian, bytes)
	default:
		return float64(0)
	}
}
