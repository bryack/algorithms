package main

import "fmt"

func calculateSpeed(fileSize float64) {
	defaultSpeed := 0.5 / 0.75
	fmt.Printf("Файл %.1f МВ загрузится за %.1f секунд\n", fileSize, fileSize/defaultSpeed)
}

func main() {
	calculateSpeed(1)
	calculateSpeed(2)
	calculateSpeed(3)
}
