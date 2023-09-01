package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
	"math/big"
	"reflect"
	"strings"
)

func RemoveProtocol(s string) string {
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "file://")
	s = strings.TrimPrefix(s, "ftp://")
	return s
}

func randint64() (int64, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}

func YoutubeId() string {
	// Generate a 64 bit number
	val, _ := randint64()

	// Encode the 64 bit number
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(val))

	return base64.RawURLEncoding.EncodeToString(b)
}

func InArray(val interface{}, array interface{}) (index bool) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
