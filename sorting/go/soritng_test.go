package sorting

import "testing"

//static UNSORTED_ARR: [i32; 10] = [5, 2, 2, 1, 4, 0, 3, 10, 24, 7];
//static SORTED_ARR: [i32; 10] = [0, 1, 2, 2, 3, 4, 5, 7, 10, 24];

var unsortedArr = [...]int{5, 2, 2, 1, 4, 0, 3, -5, 10, 24, 7}
var sortedArr = [...]int{-5, 0, 1, 2, 2, 3, 4, 5, 7, 10, 24}

func TestInsertionSort(t *testing.T) {
	newSortedArr := SortInsertion(unsortedArr[:])

	if !testArrEq(newSortedArr, sortedArr[:]) {
		t.Fatalf("Newly sorted arr (%v) doesn't equal expected (%v) arr", newSortedArr, sortedArr)
	}
}

func testArrEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
