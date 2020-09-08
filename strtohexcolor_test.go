package strtohexcolor

import "testing"

func TestXxx(t *testing.T) {
	hexColor := GenerateColor("LetsWriteStableCode")
	if hexColor != "#0FF8AF" {
		t.Errorf("Abs(-1) = %s; want #0FF8AF", hexColor)
	}
}
