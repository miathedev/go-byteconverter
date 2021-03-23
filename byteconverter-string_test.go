package byteconverter

import "testing"

func TestRead_string(t *testing.T) {
	payload := []byte{'H', 'e', 'l', 'l', 'o'}
	if Read_string(BigEndian, payload) != "Hello" {
		t.Fatalf("BigEndian String wasnt converted correct")
	}

	payload = []byte{72, 101, 108, 108, 111} //ASCII DEC
	if Read_string(BigEndian, payload) != "Hello" {
		t.Fatalf("BigEndian String wasnt converted correct")
	}

	payload = []byte{'o', 'l', 'l', 'e', 'H'}
	if Read_string(LittleEndian, payload) != "Hello" {
		t.Fatalf("LittleEndian String wasnt converted correct")
	}

}
