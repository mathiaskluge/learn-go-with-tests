package iterations

func Repeat(sequence string, iterations int) string {
	var repeated string

	for i := 0; i < iterations; i++ {
		repeated += sequence
	}
	return repeated
}
