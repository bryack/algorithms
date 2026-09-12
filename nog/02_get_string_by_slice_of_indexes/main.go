package main

// GetStringBySliceOfIndexes returns a string formed by concatenating specific characters from the input string based
// on the provided indexes.
func GetStringBySliceOfIndexes(str string, indexes []int) string {

}

func GetCharByIndex(str string, idx int) rune {
	var counter int

	for _, r := range str {
		if counter == idx {
			return r
		}

		counter++
	}

	panic("unreachable")
}
