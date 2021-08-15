package main

// Given an array of integers, return a new array such that each element at
//index i of the new array is the product of all the numbers in the original
//array except the one at i.
//
//For example, if our input was [1, 2, 3, 4, 5], the expected output would
//be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output
//would be [2, 3, 6].
//
//Follow-up: what if you can't use division?

func GetProductOfSliceWithDivision(sample []int) []int {
	// Get Products of all Elements then divide to get result
	//Solve Time (15 minutes)

	if len(sample) < 2 {
		return []int{}
	}

	response := make([]int, len(sample))
	Product := GetProductOfAllElements(sample)

	for i, factor := range sample {
		response[i] = Product / factor
	}

	return response
}

func GetProductOfAllElements(sample []int) int {
	if len(sample) < 1 {
		return -1
	}

	ProductResult := 1
	for _, factor := range sample {
		ProductResult = ProductResult * factor
	}
	return ProductResult
}

func GetProductOfSliceWithOutDivision(sample []int) []int {
	if len(sample) < 2 {
		return []int{}
	}

	response := make([]int, len(sample))
	for x, _ := range sample {
		ProductResult := 1
		for i, factor := range sample {
			if i == x {
				continue
			}
			ProductResult = ProductResult * factor
		}
		response[x] = ProductResult
	}

	return response
}
