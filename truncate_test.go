package truncate

import (
	"fmt"
	"strings"
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
					"test", 3, "…st"},
				{"works with utf8 strings",
					"тест", 3, "…ст"},
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
					"testttest", 5, "te…st"},
				{"works with loner strings even cut",
					"testttest", 4, "t…st"},
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

func TestTruncate(t *testing.T) {
	type args struct {
		str      string
		length   int
		omission string
		pos      TruncatePosition
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"omission shorter than length",
			args{"test string", 2, "...", PositionEnd},
			"te",
		},
		{
			"negative length",
			args{"test string", -2, "...", PositionEnd},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Truncate should not panic with `%v`", r)
				}
			}()

			if got := Truncate(tt.args.str, tt.args.length, tt.args.omission, tt.args.pos); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzTruncate(f *testing.F) {
	f.Add("test", 4, "…", uint8(0))
	f.Fuzz(func(t *testing.T, str string, length int, omission string, posRaw uint8) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Truncate should not panic with `%v`", r)
			}
		}()
		pos := posRaw % 3
		Truncate(str, length, omission, TruncatePosition(pos))
	})
}

func BenchmarkTruncate(b *testing.B) {
	benchmarks := []struct {
		name     string
		position TruncatePosition
	}{
		{"PositionEnd", PositionEnd},
		{"PositionStart", PositionStart},
		{"PositionMiddle", PositionMiddle},
	}
	var cases = make([]string, 100)
	for i := 0; i < 100; i++ {
		cases[i] = strings.Repeat("Ю", i)
	}

	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := range cases {
					Truncate(cases[i], i+1/2, "Я", bench.position)
				}
			}
		})
	}
}
