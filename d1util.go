// Package d1util is a library of low-level utilities for Dofus 1, reverse engineered from its client (originally written in ActionScript 2).
package d1util

import (
	"net/url"
	"strconv"
)

const (
	cellWidth      = 53.0
	cellHeight     = 27.0
	cellHalfWidth  = 26.5
	cellHalfHeight = 13.5
	levelHeight    = 20.0
)

func DecipherGameMap(data, key string) (string, error) {
	preparedKey, err := prepareKey(key)
	if err != nil {
		return "", err
	}

	checksum := checksumAlt(preparedKey)

	deciphered, err := decipherData(data, preparedKey, checksum*2)
	if err != nil {
		return "", err
	}

	return deciphered, nil
}

func unescape(s string) (string, error) {
	d, err := url.PathUnescape(s)
	if err != nil {
		return "", err
	}
	return d, nil
}

func escape(str []byte) []byte {
	escapedString := ""
	for _, c := range str {
		switch c {
		case '+':
			escapedString += "%2B"
		case '%':
			escapedString += "%25"
		default:
			escapedString += string(c)
		}
	}
	return []byte(escapedString)
}

func decodeBase16(base16 string) []byte {
	var decoded []byte
	for i := 0; i < len(base16); i += 2 {
		v, _ := strconv.ParseInt(base16[i:i+2], 16, 8)
		decoded = append(decoded, byte(v))
	}
	return decoded
}
