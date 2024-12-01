package char

import (
	"fmt"
	"time"
)

func ConvertToChars(d time.Duration, table [][]string, chars [][]string) [][]string {
	if len(table) < 10 {
		panic("incorrect table with wrong length")
	}

	hours := d / time.Hour
	chars = appendNumber(int(hours), table, chars)
	d -= hours * time.Hour

	minutes := d / time.Minute
	chars = appendNumber(int(minutes), table, chars)
	d -= minutes * time.Minute

	seconds := d / time.Second
	chars = appendNumber(int(seconds), table, chars)
	return chars
}

func appendNumber(num int, table, chars [][]string) [][]string {
	switch {
	case num < 10:
		return append(chars, table[0], table[num])
	case num < 100:
		return append(chars, table[num/10], table[num%10])
	default:
		panic(fmt.Errorf("num should >= 0 and < 100, got: %d", num))
	}
}
