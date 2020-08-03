package algorithms

/*
	bases above 10 are represented with letters (eg A, B, C - think of hex numbers)


*/

func convertDecToBase(dec int, base int) string {

	const charset = "0123456789ABCDEF"
	var result string
	for dec > 0 {
		remainder := dec % base
		dec = dec / base
		result = string(charset[remainder]) + result // if base is >= 10, use letters
	}

	return result
}
