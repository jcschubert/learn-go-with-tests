package romanLiterals

import (
	"fmt"
	"testing"
)

// Strings for formatting test messages
const (
	convertArabicToRoman = "%d gets converted to %q"
	convertArabicToRomanFailed = "Conversion result: %q, expected %q"
	convertRomanToArabic = "%q gets converted to %d"
	convertRomanToArabicFailed = "Conversion result: %d, expected %d"

)

type RomanLiteralCases struct {
	Arabic int
	Roman  string
}

var cases []RomanLiteralCases = []RomanLiteralCases{
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

func TestRomanNumerals(t *testing.T) {
	runCases(t, cases, ConvertToRoman)
}

func runCases(t *testing.T, cases []RomanLiteralCases, f func(int) string) {
	for _, test := range cases {
		t.Run(fmt.Sprintf(convertArabicToRoman, test.Arabic, test.Roman), func(t *testing.T) {
			got := f(test.Arabic)
			if got != test.Roman {
				t.Errorf(convertArabicToRomanFailed, got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases[:2] {
		t.Run(fmt.Sprintf(convertRomanToArabic, test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf(convertRomanToArabicFailed, got, test.Arabic)
			}
		})
	}
}
