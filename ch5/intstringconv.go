package ch5

import (
	"errors"
	"math"
)

var errSyntax = errors.New("ch5.stringToInt: invalid syntax")
var errRange = errors.New("ch5.stringToInt: value out of range")

// StringToInt converts number represented by string with base 10 to integer.
func StringToInt(s string) (int64, error) {
	const cutoff = math.MaxInt64/10 + 1 // The first smallest number such that cutoff*10 > MaxInt64.

	if len(s) == 0 {
		return 0, errSyntax
	}

	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	var u uint64
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0, errSyntax
		}

		if u >= cutoff {
			// u*10 overflows.
			return 0, errRange
		}
		u *= 10

		nu := u + uint64(s[i]-'0')
		if neg && nu > -math.MinInt64 || !neg && nu > math.MaxInt64 {
			// u+v overflows.
			return 0, errRange
		}
		u = nu
	}

	n := int64(u)
	if neg {
		n = -n
	}
	return n, nil
}

// IntToString converts integer to string.
func IntToString(n int64) string {
	if n == 0 {
		return "0"
	}

	d := 0 // Number of digits in result string.
	neg := false
	u := uint64(n)
	if n < 0 {
		d++
		neg = true
		u = uint64(^n + 1)
	}

	d += 1 + int(math.Log10(float64(u))) // Number of digits of u.
	s := make([]byte, d)
	mi := d - 1
	for u > 0 {
		s[mi] = byte(u%10 + '0')
		u /= 10
		mi--
	}

	if neg {
		s[0] = '-'
	}
	return string(s)
}