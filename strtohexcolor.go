package strtohexcolor

import (
	"strconv"
	"strings"
)

// GenerateColor generate uniquely color based on given str
// Web colors are colors used in displaying web pages on the World Wide Web
// and the methods for describing and specifying those colors
// Colors may be specified as an RGB triplet or in hexadecimal format (a hex triplet) or
// according to their common English names in some cases
// A color tool or other graphics software is often used to generate color values
// In some uses, hexadecimal color codes are specified with notation using a leading number sign (#).[1][2]
// A color is specified according to the intensity of its red, green and blue components, each represented by eight bits. Thus,
// there are 24 bits used to specify a web color within the sRGB gamut, and 16,777,216 colors that may be so specified.
// copied from [this](https://stackoverflow.com/a/16348977/13961710) 🤪
func GenerateColor(str string) string {
	hash := 0

	for i := 0; i < len(str); i++ {
		hash = int([]rune(string(str[i]))[0]) + ((hash << 5) - hash)
	}

	color := "#"

	for i := 0; i < 3; i++ {
		value := (hash >> (i * 8)) & 0xFF
		hex := ("00" + strconv.FormatInt(int64(value), 16))
		color += hex[len(hex)-2:]
	}

	return strings.ToUpper(color)
}
