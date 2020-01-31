package romanLiterals

import(
	"testing"
	"fmt"
)

// Strings for formatting test messages
const(
	msgConversionDescription = "%d gets converted to %q"
	msgConversionFailed = "Conversion result: %q, expected: %q"
)

type RomanLiteralCases struct {
	Arabic int
	Roman  string
}

func TestRomanNumerals(t *testing.T) {
	cases := []RomanLiteralCases{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}

	runCases(t, cases, ConvertToRoman)
}

func runCases(t *testing.T, cases []RomanLiteralCases, f func(int) string) {
	for _, test := range cases {
		t.Run(fmt.Sprintf(msgConversionDescription, test.Arabic, test.Roman), func(t *testing.T) {
			got := f(test.Arabic)
			if got != test.Roman {
				t.Errorf(msgConversionFailed, got, test.Roman)
			}
		})
	}
}
