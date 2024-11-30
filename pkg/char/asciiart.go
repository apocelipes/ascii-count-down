package char

import "time"

var ASCIIArtChars = [][]string{
	{
		`  ___  `,
		` / _ \ `,
		`| | | |`,
		`| | | |`,
		`| |_| |`,
		` \___/ `,
	},
	{
		` __ `,
		`/_ |`,
		` | |`,
		` | |`,
		` | |`,
		` |_|`,
	},
	{
		` ___  `,
		`|__ \ `,
		`   ) |`,
		`  / / `,
		` / /_ `,
		`|____|`,
	},
	{
		` ____  `,
		`|___ \ `,
		`  __) |`,
		` |__ < `,
		` ___) |`,
		`|____/ `,
	},
	{
		` _  _   `,
		`| || |  `,
		`| || |_ `,
		`|__   _|`,
		`   | |  `,
		`   |_|  `,
	},
	{
		` _____ `,
		`| ____|`,
		`| |__  `,
		`|___ \ `,
		` ___) |`,
		`|____/ `,
	},
	{
		`   __  `,
		`  / /  `,
		` / /_  `,
		`| '_ \ `,
		`| (_) |`,
		` \___/ `,
	},
	{
		` ______ `,
		`|____  |`,
		`    / / `,
		`   / /  `,
		`  / /   `,
		` /_/    `,
	},
	{
		`  ___  `,
		` / _ \ `,
		`| (_) |`,
		` > _ < `,
		`| (_) |`,
		` \___/ `,
	},
	{
		`  ___  `,
		` / _ \ `,
		`| (_) |`,
		` \__, |`,
		`   / / `,
		`  /_/  `,
	},
	{
		`   `,
		` _ `,
		`(_)`,
		`   `,
		` _ `,
		`(_)`,
	},
}

const ASCIIArtColonIdx = 10

func MaxASCIIArtCharWidth() int {
	width := len(ASCIIArtChars[0][0])
	for _, digit := range ASCIIArtChars {
		width = max(width, len(digit[0]))
	}
	return width
}

func MaxASCIIArtCharHeight() int {
	return len(ASCIIArtChars[0])
}

func ConvertToASCIIArtChars(d time.Duration, chars [][]string) [][]string {
	hour := d / time.Hour
	if hour < 10 {
		chars = append(chars, ASCIIArtChars[0], ASCIIArtChars[hour])
	} else {
		chars = append(chars, ASCIIArtChars[hour/10], ASCIIArtChars[hour%10])
	}
	d -= hour * time.Hour
	minute := d / time.Minute
	if minute < 10 {
		chars = append(chars, ASCIIArtChars[0], ASCIIArtChars[minute])
	} else {
		chars = append(chars, ASCIIArtChars[minute/10], ASCIIArtChars[minute%10])
	}
	d -= minute * time.Minute
	second := d / time.Second
	if second < 10 {
		chars = append(chars, ASCIIArtChars[0], ASCIIArtChars[second])
	} else {
		chars = append(chars, ASCIIArtChars[second/10], ASCIIArtChars[second%10])
	}
	return chars
}
