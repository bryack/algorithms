package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])

		if lo <= hi {
			res = append(res, []int{lo, hi})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
