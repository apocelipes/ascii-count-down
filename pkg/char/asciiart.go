package char

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

func MaxASCIIArtCharWidth(until int) int {
	if until > 9 || until < 0 {
		return 0
	}
	width := len(ASCIIArtChars[0][0])
	for _, digit := range ASCIIArtChars[:until+1] {
		width = max(width, len(digit[0]))
	}
	return width
}

func MaxASCIIArtCharHeight() int {
	return len(ASCIIArtChars[0])
}
