package truncate

import (
	"fmt"
	"testing"
)

func TestTruncator(t *testing.T) {
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
			CutStrategy{},
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
			CutEllipsisStrategy{},
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
		{
			CutEllipsisLeadingStrategy{},
			[]tcase{
				{"works with shorter strings",
					"те", 10, "те"},
				{"works with exact size strings",
					"тест", 4, "тест"},
				{"works with ansi strings",
					"test", 3, "…te"},
				{"works with utf8 strings",
					"тест", 3, "…те"},
			},
		},
		{
			EllipsisMiddleStrategy{},
			[]tcase{
				{"works with shorter strings",
					"те", 10, "те"},
				{"works with exact size strings",
					"тест", 4, "тест"},
				{"works with ansi strings",
					"test", 3, "t…t"},
				{"works with utf8 strings",
					"тест", 3, "т…т"},
				{"works with loner strings off cut",
					"testttest", 5, "te…tt"},
				{"works with loner strings even cut",
					"testttest", 4, "te…t"},
			},
		},
	}
	for _, tt := range tests {
		for _, cc := range tt.cases {
			t.Run(fmt.Sprintf("%T - %s", tt.strategy, cc.name), func(t *testing.T) {
				if got := Truncator(cc.str, cc.length, tt.strategy); got != cc.want {
					t.Errorf("Truncate(`%v`) = `%v`, want `%v`", cc.str, got, cc.want)
				}
			})
		}
	}
}
