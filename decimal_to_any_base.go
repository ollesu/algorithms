package algorithms

func convertDecToBase(dec int, base int) string {
	var result string
	for dec > 0 {
		remainder := dec % base
		dec = dec / base
		result = string(remainder) + result
	}

	return result
}
