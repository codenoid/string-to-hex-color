package strtohexcolor

import "testing"

func TestXxx(t *testing.T) {
	hexColor := GenerateColor("Let's write Stable Software")
	if hexColor != "#B25320" {
		t.Errorf("Abs(-1) = %s; want #B25320", hexColor)
	}
}
