package date

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	yearArr := [2][12]int{
		{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
		{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	}

	dateArr := strings.Split(date, "-")
	year, err := strconv.Atoi(dateArr[0])
	if err != nil {
		return -1
	}

	index := 0
	if year % 4 == 0 {
		index = 1
	}

	month, _ := strconv.Atoi(dateArr[1])
	day, _ := strconv.Atoi(dateArr[2])

	ans := 0
	for i := 0; i < month; i++ {
		ans += yearArr[index][i]
	}

	return ans + day
}
