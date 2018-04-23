package lib

import (
	"crypto/rand"
	"encoding/base64"
)

func genByte(i int) ([]byte, error) {
	var b = make([]byte, i)
	_,err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RandomString(size int) (string, error) {
	b, err := genByte(size)
	return base64.URLEncoding.EncodeToString(b), err
}
