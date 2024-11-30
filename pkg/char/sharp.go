package char

import "time"

var SharpChars = [][]string{
	{
		" ####   ",
		" #    # ",
		"#      #",
		"#      #",
		"#      #",
		" #    # ",
		"  ####  ",
	},
	{
		"    #   ",
		"   ##   ",
		"  # #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		" #######",
	},
	{
		"######  ",
		"#      #",
		"       #",
		" ###### ",
		"#       ",
		"#       ",
		"########",
	},
	{
		" #####  ",
		" #     #",
		"       #",
		"  ##### ",
		"       #",
		" #     #",
		"  ##### ",
	},
	{
		"#       ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #######",
		"      # ",
		"      # ",
	},
	{
		"####### ",
		" #      ",
		" #      ",
		" ###### ",
		"       #",
		" #     #",
		"  ##### ",
	},
	{
		"######  ",
		"#      #",
		"#       ",
		"####### ",
		"#      #",
		"#      #",
		" ###### ",
	},
	{
		"########",
		" #     #",
		"      # ",
		"     #  ",
		"    #   ",
		"    #   ",
		"    #   ",
	},
	{
		"######  ",
		"#      #",
		"#      #",
		" ###### ",
		"#      #",
		"#      #",
		" ###### ",
	},
	{
		"######  ",
		"#      #",
		"#      #",
		" #######",
		"       #",
		"#      #",
		" ###### ",
	},
	{
		"  ",
		"##",
		"  ",
		"  ",
		"  ",
		"##",
		"  ",
	},
}

const SharpColonIdx = 10

func MaxSharpCharWidth() int {
	return len(SharpChars[0][0])
}

func MaxSharpCharHeight() int {
	return len(SharpChars[0])
}

func ConvertToSharpChars(d time.Duration, chars [][]string) [][]string {
	hour := d / time.Hour
	if hour < 10 {
		chars = append(chars, SharpChars[0], SharpChars[hour])
	} else {
		chars = append(chars, SharpChars[hour/10], SharpChars[hour%10])
	}
	d -= hour * time.Hour
	minute := d / time.Minute
	if minute < 10 {
		chars = append(chars, SharpChars[0], SharpChars[minute])
	} else {
		chars = append(chars, SharpChars[minute/10], SharpChars[minute%10])
	}
	d -= minute * time.Minute
	second := d / time.Second
	if second < 10 {
		chars = append(chars, SharpChars[0], SharpChars[second])
	} else {
		chars = append(chars, SharpChars[second/10], SharpChars[second%10])
	}
	return chars
}
