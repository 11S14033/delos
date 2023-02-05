package delos

func question2(totalStudent, totalCandies, firstStudent int64) int64 {
	var lastStudent int64
	for totalCandies > 0 {
		lastStudent = firstStudent

		if firstStudent == totalStudent {
			firstStudent = 0
		}

		firstStudent++
		totalCandies--
	}

	return lastStudent
}
