package base62

import (
	"fmt"
	"strings"
)

// All characters
const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = int64(len(alphabet))
)
func TransTo62(id int64)string{
	// 1 -- > 1
	// 10-- > a
	// 61-- > Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for{
		var result byte
		number := id % 62
		result = charset[number]
		var tmp []byte
		tmp = append(tmp,result)
		shortUrl = append(tmp,shortUrl...)
		id = id / 62
		if id == 0{
			break
		}
	}
	fmt.Println(string(shortUrl))
	return string(shortUrl)
}

// Encode number to base62.
func Encode(n int64) string {
	if n == 0 {
		return string(alphabet[0])
	}

	s := ""
	for ; n > 0; n = n / length {
		s = string(alphabet[n%length]) + s
	}
	return s
}

// Decode converts a base62 token to int.
func Decode(key string) (int64, error) {
	var n int64
	for _, c := range []byte(key) {
		i := strings.IndexByte(alphabet, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		n = length*n + int64(i)
	}
	return n, nil
}
