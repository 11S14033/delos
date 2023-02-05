package delos

func question3(arr []int32) string {
	var right, left int32
	for _, key := range arr {
		right += key
	}

	for _, key := range arr {
		right = right - key
		if right == left {
			return "YES"
		}

		left += key
	}

	return "NO"
}
