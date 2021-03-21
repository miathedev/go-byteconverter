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

import (
	"fmt"
	"testing"
)

//Tools:
//Converter tool: https://gregstoll.com/~gregstoll/floattohex/
//http://www.binaryconvert.com/result_signed_int.html?decimal=051050054052053051053051053
func TestRead_float32(t *testing.T) {
	bytes := []byte{0xb1, 0xe1, 0xb5, 0x41} //22.7352
	mynumber := Read_float32(LittleEndian, bytes)
	s := fmt.Sprintf("%f", mynumber)

	if mynumber != float32(22.7352) {
		t.Fatalf("float32 failure")
	}
	fmt.Println(s)

	bytes = []byte{0x41, 0xb5, 0xe1, 0xb1} //22.7352
	mynumber = Read_float32(BigEndian, bytes)
	s = fmt.Sprintf("%f", mynumber)

	if mynumber != float32(22.7352) {
		t.Fatalf("float32 failure")
	}
	fmt.Println(s)
}

func TestRead_uint8(t *testing.T) {
	bytes := []byte{21}
	mynumber := Read_uint8(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 21 {
		t.Fatalf("uint8 failure")
	}
	fmt.Println(s)

	bytes = []byte{21}
	mynumber = Read_uint8(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != 21 {
		t.Fatalf("uint8 failure")
	}
	fmt.Println(s)
}

func TestRead_uint16(t *testing.T) {
	bytes := []byte{0x02, 0x52} //594
	mynumber := Read_uint16(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594 {
		t.Fatalf("uint16 failure")
	}
	fmt.Println(s)

	bytes = []byte{0x52, 0x02} //594
	mynumber = Read_uint16(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != 594 {
		t.Fatalf("uint16 failure")
	}
	fmt.Println(s)
}
func TestRead_uint32(t *testing.T) {
	bytes := []byte{0x23, 0x6F, 0xE3, 0xA7} //594535335
	mynumber := Read_uint32(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594535335 {
		t.Fatalf("uint32 failure")
	}
	fmt.Println(s)

	bytes = []byte{0xA7, 0xE3, 0x6F, 0x23} //594535335
	mynumber = Read_uint32(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != 594535335 {
		t.Fatalf("uint32 failure")
	}
	fmt.Println(s)
}

func TestRead_int8(t *testing.T) {
	bytes := []byte{0x17} //23
	mynumber := Read_int8(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 23 {
		t.Fatalf("int8 failure")
	}
	fmt.Println(s)

	bytes = []byte{0x17} //23
	mynumber = Read_int8(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != 23 {
		t.Fatalf("int8 failure")
	}
	fmt.Println(s)
}

func TestRead_int16(t *testing.T) {
	bytes := []byte{0xFD, 0x72} //-654
	mynumber := Read_int16(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != -654 {
		t.Fatalf("int16 failure")
	}
	fmt.Println(s)

	bytes = []byte{0x72, 0xFD} //-654
	mynumber = Read_int16(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != -654 {
		t.Fatalf("int16 failure")
	}
	fmt.Println(s)
}

func TestRead_int32(t *testing.T) {
	bytes := []byte{0xD8, 0xFC, 0x6D, 0x75} //-654545547
	mynumber := Read_int32(BigEndian, bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != -654545547 {
		t.Fatalf("int32 failure")
	}
	fmt.Println(s)

	bytes = []byte{0x75, 0x6D, 0xFC, 0xD8} //-654
	mynumber = Read_int32(LittleEndian, bytes)
	s = fmt.Sprintf("%d", mynumber)

	if mynumber != -654545547 {
		t.Fatalf("int32 failure")
	}
	fmt.Println(s)
}
