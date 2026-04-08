package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var builder strings.Builder
	builder.Grow(15)

	for i := 0; i < len(allRomanNumerals); i++ {

		for num >= allRomanNumerals[i].Value {
			builder.WriteString(allRomanNumerals[i].Symbol)
			num = num - allRomanNumerals[i].Value
		}
	}

	return builder.String()
}
