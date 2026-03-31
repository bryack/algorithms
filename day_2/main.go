package main

import "fmt"

func progress(total, process int) {
	res := float64(process) / float64(total)
	fmt.Printf("Обработано %d из %d (%.1f %%)\n", process, total, res*100)
}

func main() {
	progress(8, 3)
	progress(8, 4)
	progress(8, 8)
}
