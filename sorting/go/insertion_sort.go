package sorting

func SortInsertion(arr []int) []int {
	sortedArr := make([]int, len(arr))

	copy(sortedArr, arr)

	for i := 0; i < len(sortedArr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if sortedArr[j-1] > sortedArr[j] {
				sortedArr[j-1], sortedArr[j] = sortedArr[j], sortedArr[j-1]
			}
		}
	}

	return sortedArr
}
