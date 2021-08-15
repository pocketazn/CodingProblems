package main

const (
	k = 17
)

func DoesListContainSumPairEqualToK(givenList []int) bool {
	// validate slice
	if len(givenList) < 2 {
		return false
	}

	CheckedValues := make(map[int]bool)

	for start, ValueX := range givenList {
		// Did we already test this value
		if _, ok := CheckedValues[ValueX]; ok {
			continue
		}

		// Test Value against negfhhtdfxt values in slice to the end
		for i := start + 1; i < len(givenList); i++ {
			if CompareSumToK(ValueX, givenList[i]) {
				return true
			}
			CheckedValues[ValueX] = true
		}
	}

	return false
}

func CompareSumToK(x int, y int) bool {
	if (x + y) == k {
		return true
	}
	return false
}
