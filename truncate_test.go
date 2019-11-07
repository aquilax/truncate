package truncate

import (
	"fmt"
	"testing"
)

func TestTruncate(t *testing.T) {
	type tcase struct {
		name   string
		str    string
		length int
		want   string
	}
	tests := []struct {
		strategy Strategy
		cases    []tcase
	}{
		{
			Cut{},
			[]tcase{
				{"works with shorter strings",
					"те", 10, "те"},
				{"works with exact size strings",
					"тест", 4, "тест"},
				{"works with ansi strings",
					"test", 3, "tes"},
				{"works with utf8 strings",
					"тест", 3, "тес"},
			},
		},
		{
			CutEllipsis{},
			[]tcase{
				{"works with shorter strings",
					"те", 10, "те"},
				{"works with exact size strings",
					"тест", 4, "тест"},
				{"works with ansi strings",
					"test", 3, "te…"},
				{"works with utf8 strings",
					"тест", 3, "те…"},
			},
		},
	}
	for _, tt := range tests {
		for _, cc := range tt.cases {
			t.Run(fmt.Sprintf("%T - %s", tt.strategy, cc.name), func(t *testing.T) {
				if got := Truncate(cc.str, cc.length, tt.strategy); got != cc.want {
					t.Errorf("Truncate() = `%v`, want `%v`", got, cc.want)
				}
			})
		}
	}
}
