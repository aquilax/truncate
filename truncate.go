// Package truncate provides set of strategies to truncate strings
package truncate

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
