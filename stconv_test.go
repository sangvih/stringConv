package stringConv

import "testing"

func TestDz(t *testing.T) {
	input := "d2a3"
	output := "ddaaa"
	funcOut, err := Dz(input)
	if funcOut != output {
		t.Error("Ошибка - ", err)
	}
}
