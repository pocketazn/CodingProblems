package _2021_10

import (
	"strconv"
	"strings"
)

func GetPerfectNumber(sample int) int {
	number := strconv.Itoa(sample)
	parsedNum := strings.Split(number, "")
	total := 0
	for _, num := range parsedNum {
		i, _ := strconv.Atoi(num)
		total = total + i
		if total > 10 {
			return -1
		}
	}

	if total != 10 {
		parsedNum = append(parsedNum, strconv.Itoa(10-total))
	}
	joinedNumbers := strings.Join(parsedNum, "")
	i, _ := strconv.Atoi(joinedNumbers)
	return i
}