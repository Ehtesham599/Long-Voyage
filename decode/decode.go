package decode

import "encoding/hex"

// Check if input hex string is valid
func IsValidHexString(str string) bool {
	if str != "" {
		return len(str)%2 == 0
	}

	return false
}

// decode string to byte slice
func StrToByteSlice(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}
