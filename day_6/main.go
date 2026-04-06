package main

import (
	"math"
)

func CelsiusFahrenheit(temp float64) float64 {
	t := temp*9/5 + 32
	return math.Round(t*10) / 10
}

func FahrenheitCelsius(temp float64) float64 {
	t := (temp - 32) * 5 / 9
	return math.Round(t*10) / 10
}

func RUBUSD(rub, kopeck int, rate float64) (int, int) {
	totalRub := float64(rub) + float64(kopeck)/100
	usd := totalRub / rate
	dollars := int(usd)
	cents := int(math.Round((usd - float64(dollars)) * 100))
	return dollars, cents
}

func USDRUB(d, c int, rate float64) (int, int) {
	totalUSD := float64(d) + float64(c)/100
	usd := totalUSD / rate
	rub := int(usd)
	kop := int(math.Round((usd - float64(rub)) * 100))
	return rub, kop
}

func convertToTitle(columnNumber int) string {
	if columnNumber > 26 {
		s := []rune{}
		for {
			remainder := columnNumber % 26
			quotient := columnNumber / 26
			if remainder == 0 {
				remainder = 26
				quotient--
			}
			s = append(s, rune(int('A')+(remainder-1)))
			if quotient == 0 {
				break
			}
			columnNumber = quotient
		}
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return string(s)
	}

	char := rune(int('A') + (columnNumber - 1))

	return string(char)
}
