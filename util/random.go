package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	num           = "1234567890"
	SingaporeCode = "(+65)"
	MalaysiaCode  = "(+60)"
)

func Random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomNumberInString(n int) string {
	r := Random()
	var sb strings.Builder
	k := len(num)

	for i := 0; i < n; i++ {
		c := num[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
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
		return concatContactNumber(SingaporeCode, RandomNumberInString(8))
	}
	return concatContactNumber(MalaysiaCode, RandomNumberInString(9))
}

func concatContactNumber(code, num string) string {
	var sb strings.Builder
	sb.WriteString(code)
	sb.WriteString(num)
	return sb.String()
}

func RandomTime() time.Time {
	return time.Now()
}
