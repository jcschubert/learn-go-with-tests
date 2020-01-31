package romanLiterals

import "testing"

type RomanLiteralCases struct {
	Description string
	Arabic      int
	Want        string
}

func TestRomanNumerals(t *testing.T) {
	cases := []RomanLiteralCases{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
	}

	testCases(t, cases, ConvertToRoman)
	/*
		for _, test := range cases {
			t.Run(test.Description, func(t *testing.T) {
				got := ConvertToRoman(test.Arabic)
				if got != test.Want {
					t.Errorf("got %q, want %q", got, test.Want)
				}
			})
		}*/
}

func testCases(t *testing.T, cases []RomanLiteralCases, f func(int) string) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := f(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
