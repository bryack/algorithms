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
