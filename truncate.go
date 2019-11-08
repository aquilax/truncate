// Package truncate provides set of strategies to truncate strings
package truncate

import "math"

// Strategy is a  interface for truncation strategy
type Strategy interface {
	Truncate(string, int) string
}

// Truncate cuts a string to length using the truncation strategy
func Truncate(str string, length int, strategy Strategy) string {
	return strategy.Truncate(str, length)
}

// CutStrategy simply truncates the string to the desired length
type CutStrategy struct{}

func (CutStrategy) Truncate(str string, length int) string {
	r := []rune(str)
	if length >= len(r) {
		return str
	}
	return string(r[0:length])
}

// CutEllipsisStrategy simply truncates the string to the desired length and adds ellipsis at the end
type CutEllipsisStrategy struct{}

func (s CutEllipsisStrategy) Truncate(str string, length int) string {
	r := []rune(str)
	if length >= len(r) {
		return str
	}
	return string(r[0:length-1]) + "…"
}

// EllipsisMiddleStrategy truncates the string to the desired length and adds ellipsis in the middle
type EllipsisMiddleStrategy struct{}

func (e EllipsisMiddleStrategy) Truncate(str string, length int) string {
	r := []rune(str)
	sLen := len(r)
	if length >= sLen {
		return str
	}
	if length < 3 {
		return CutStrategy{}.Truncate(str, length)
	}
	var delta int
	if sLen%2 == 0 {
		delta = int(math.Ceil(float64(sLen-length) / 2))
	} else {
		delta = int(math.Floor(float64(sLen-length) / 2))
	}
	result := make([]rune, length)
	copy(result, r[0:delta])
	result[delta] = '…'
	copy(result[delta+1:], r[length-delta+1:])
	return string(result)
}
