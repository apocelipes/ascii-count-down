package util

import "flag"

func StringFlagWithShortName(longName, shortName, defaultValue, usage string) *string {
	if longName == "" || shortName == "" {
		panic("flag name should not be empty")
	}
	p := flag.String(longName, defaultValue, usage)
	flag.StringVar(p, shortName, defaultValue, usage)

	return p
}
