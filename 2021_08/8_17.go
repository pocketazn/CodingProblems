package _2021_08

import "sort"

// Given an array of integers, find the first missing
//positive integer in linear time and constant space.
//In other words, find the lowest positive integer that
//does not exist in the array. The array can contain duplicates
//and negative numbers as well.
//
//For example, the input [3, 4, -1, 1] should give 2.
//The input [1, 2, 0] should give 3.
//
//You can modify the input array in-place.

// time to solve 20 mins

func FindLowestPositiveInt(sample []int) int {
	if len(sample) == 0 {
		return 0
	}

	sort.Ints(sample)

	if sample[len(sample)-1] < 0 {
		return 0
	}

	var lowestPositiveUnit *int
	for i, unit := range sample {
		if unit < 0 || (i > 0 && unit == sample[i-1]) {
			continue
		}

		if unit >= 0 && lowestPositiveUnit == nil {
			firstPositiveUnit := unit
			lowestPositiveUnit = &firstPositiveUnit
		}

		if *lowestPositiveUnit != unit {
			return *lowestPositiveUnit
		} else if lowestPositiveUnit != nil {
			*lowestPositiveUnit++
		}

		if i == len(sample)-1 {
			if lowestPositiveUnit == nil {
				return 0
			}
		}
	}

	return sample[len(sample)-1] + 1

}
