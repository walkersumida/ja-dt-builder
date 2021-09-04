package jadtbuilder

import (
	"testing"
)

func TestBuild(t *testing.T) {
	expected := "9月4日(土) 13:00"
	got := Build("2021 9 4 13")

	if got != expected {
		t.Errorf("\nexpected: %s\n     got: %s", expected, got)
	}
}
