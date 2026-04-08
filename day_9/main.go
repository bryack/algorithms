package main

import (
	"slices"
	"strings"
)

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

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	s := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			s += prices[i+1] - prices[i]
		}
	}

	return s
}

func lemonadeChange(bills []int) bool {
	ten, five := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five < 1 {
				return false
			}
			ten++
			five--
		case 20:
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	slices.Sort(g)
	slices.Sort(s)

	child, cookie := 0, 0

	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			child++
		}
		cookie++
	}
	return child
}
