package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	num           = "1234567890"
	SingaporeCode = "(+65)"
	MalaysiaCode  = "(+60)"
)

func Random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateString(content string, limit int) string {
	r := Random()
	var sb strings.Builder
	k := len(content)

	for i := 0; i < limit; i++ {
		c := content[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomNumberInString(n int) string {
	return generateString(num, n)
}

func RandomString(n int) string {
	return generateString(alphabet, n)
}

func RandomCountryCode() string {
	r := Random()
	countryCode := []string{SingaporeCode, MalaysiaCode}
	k := len(countryCode)

	return countryCode[r.Intn(k)]
}

func RandomUsername() string {
	countryCode := RandomCountryCode()
	if countryCode == SingaporeCode {
		return concatTwoStrings(SingaporeCode, RandomNumberInString(8))
	}
	return concatTwoStrings(MalaysiaCode, RandomNumberInString(9))
}

func RandomCarPlate() string {
	return concatTwoStrings(RandomString(3), RandomNumberInString(4))
}

func concatTwoStrings(code, num string) string {
	var sb strings.Builder
	sb.WriteString(code)
	sb.WriteString(num)
	return sb.String()
}

func RandomTime() time.Time {
	return time.Now()
}
