package utils

var limit = 10

func CalcOffset(p int) int {
	return p * limit
}

func ConcatArgs(args []string) string {
	q := ""
	for _, arg := range args {
		q += arg + " "
	}
	return q
}
