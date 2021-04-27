package d1util

import (
	"encoding/hex"
	"math"
)

var hexChars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func prepareKey(key string) (string, error) {
	b, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	s, err := unescape(string(b))
	if err != nil {
		return "", err
	}

	return s, nil
}

func decipherData(data, key string, c int) (string, error) {
	rawData := decodeBase16(data)

	decryptedData := make([]byte, len(rawData))
	for i := 0; i < len(rawData); i++ {
		decryptedData[i] = rawData[i] ^ key[(c+i)%len(key)]
	}
	return string(decryptedData), nil
}

func checksum(data string) string {
	sum := 0
	for _, v := range data {
		sum += int(v) % 16
	}
	return hexChars[sum%16]
}

func checksumAlt(data string) int {
	sum := 0
	for _, v := range data {
		sum += int(v) % 16
	}
	return sum % 16
}

func d2h(d int) string {
	if d > 255 {
		d = 255
	}
	return hexChars[int(math.Floor(float64(d)/16))] + hexChars[d%16]
}
