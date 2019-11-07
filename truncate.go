// Package truncate provides set of strategies to truncate strings
package truncate

// Strategy is a  interface for truncation strategy
type Strategy interface {
	truncate(string, int) string
}

// Truncate cuts a string to length using the truncation strategy
func Truncate(str string, length int, strategy Strategy) string {
	return strategy.truncate(str, length)
}

// Cut simply truncates the string to the desired length
type Cut struct{}

func (c Cut) truncate(str string, length int) string {
	r := []rune(str)
	if length >= len(r) {
		return str
	}
	return string(r[0:length])
}

// CutEllipsis simply truncates the string to the desired length and adds ellipsis at the end
type CutEllipsis struct{}

func (c CutEllipsis) truncate(str string, length int) string {
	r := []rune(str)
	if length >= len(r) {
		return str
	}
	return string(r[0:length-1]) + "…"
}
