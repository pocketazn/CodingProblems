package _2021_08

const (
	k = 17
)

//Given a list of numbers and a number k, return whether any two numbers
//from the list add up to k.
//
//For example, given [10, 15, 3, 7] and k of 17, return true
//since 10 + 7 is 17.
//
//Bonus: Can you do this in one pass?

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

		// Test Value against next values in slice to the end
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
