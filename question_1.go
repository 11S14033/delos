package delos

type schedule struct {
	d, m, y int
}

func question1(expected, returned schedule) int {
	if expected.y > returned.y {
		return 0
	}

	dailyFine := 15
	monthlyFine := 500
	annualFine := 12000

	if expected.y < returned.y {
		return annualFine * (returned.y - expected.y)
	} else {
		if expected.m < returned.m {
			return monthlyFine * (returned.m - expected.m)
		} else if expected.d < returned.d && expected.m <= returned.m {
			return dailyFine * (returned.d - expected.d)
		}

	}

	return 0
}
