package util

func Concat(input []string, separator string) string {
	concat := ""
	for _, pre := range input {
		if concat == "" {
			concat += pre
		} else {
			concat += separator + pre
		}
	}

	return concat
}
