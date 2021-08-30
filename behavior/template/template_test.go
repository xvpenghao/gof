package template

import "testing"

func TestTemplate(t *testing.T) {
	var tomato Cooker = &Tomato{}

	DoCook(tomato)
}
