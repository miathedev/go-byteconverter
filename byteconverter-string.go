package byteconverter

func reverseBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reverseBytes(input[1:]), input[0])
}

//Converts byte array representation of a string back into a string.
func Read_string(endian int, bytes []byte) string {
	if endian != BigEndian {
		bytes = reverseBytes(bytes)
	}
	return string(bytes)
}
