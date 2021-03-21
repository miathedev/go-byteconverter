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

func TestLEUint8(t *testing.T) {
	bytes := []byte{21}
	mynumber := Read_le_uint8(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 21 {
		t.Fatalf("uint8 failure")
	}
	fmt.Println(s)
}

func TestLEUint16(t *testing.T) {
	bytes := []byte{0x52, 0x02} //594
	mynumber := Read_le_uint16(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594 {
		t.Fatalf("uint16 failure")
	}
	fmt.Println(s)
}

func TestLEUint32(t *testing.T) {
	bytes := []byte{0xA7, 0xE3, 0x6F, 0x23} //594535335
	mynumber := Read_le_uint32(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594535335 {
		t.Fatalf("uint32 failure")
	}
	fmt.Println(s)
}

func TestLEint8(t *testing.T) {
	bytes := []byte{0x17} //23
	mynumber := Read_le_int8(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 23 {
		t.Fatalf("int8 failure")
	}
	fmt.Println(s)
}

func TestLEint16(t *testing.T) {
	bytes := []byte{0x85, 0x7F} //32645
	mynumber := Read_le_int16(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 32645 {
		t.Fatalf("int16 failure")
	}
	fmt.Println(s)
}

func TestLEint32(t *testing.T) {
	bytes := []byte{0x1F, 0x49, 0x75, 0x13} //326453535
	mynumber := Read_le_int32(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 326453535 {
		t.Fatalf("int32 failure")
	}
	fmt.Println(s)
}

func TestLEfloat32(t *testing.T) {
	bytes := []byte{0xb1, 0xe1, 0xb5, 0x41} //22.7352
	mynumber := Read_le_float32(bytes)
	s := fmt.Sprintf("%f", mynumber)

	if mynumber != float32(22.7352) {
		t.Fatalf("float32 failure")
	}
	fmt.Println(s)
}

//Big endian

func TestBEUint8(t *testing.T) {
	bytes := []byte{21}
	mynumber := Read_be_uint8(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 21 {
		t.Fatalf("uint8 failure")
	}
	fmt.Println(s)
}

func TestBEUint16(t *testing.T) {
	bytes := []byte{0x02, 0x52} //594
	mynumber := Read_be_uint16(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594 {
		t.Fatalf("uint16 failure")
	}
	fmt.Println(s)
}

func TestBEUint32(t *testing.T) {
	bytes := []byte{0x23, 0x6F, 0xE3, 0xA7} //594535335
	mynumber := Read_be_uint32(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 594535335 {
		t.Fatalf("uint32 failure")
	}
	fmt.Println(s)
}

func TestBEint8(t *testing.T) {
	bytes := []byte{0x17} //23
	mynumber := Read_be_int8(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 23 {
		t.Fatalf("int8 failure")
	}
	fmt.Println(s)
}

func TestBEint16(t *testing.T) {
	bytes := []byte{0x7F, 0x85} //32645
	mynumber := Read_be_int16(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 32645 {
		t.Fatalf("int16 failure")
	}
	fmt.Println(s)
}

func TestBEint32(t *testing.T) {
	bytes := []byte{0x13, 0x75, 0x49, 0x1F} //326453535
	mynumber := Read_be_int32(bytes)
	s := fmt.Sprintf("%d", mynumber)

	if mynumber != 326453535 {
		t.Fatalf("int32 failure")
	}
	fmt.Println(s)
}

func TestBEfloat32(t *testing.T) {
	bytes := []byte{0x41, 0xb5, 0xe1, 0xb1} //22.7352
	mynumber := Read_be_float32(bytes)
	s := fmt.Sprintf("%f", mynumber)

	if mynumber != float32(22.7352) {
		t.Fatalf("float32 failure")
	}
	fmt.Println(s)
}
