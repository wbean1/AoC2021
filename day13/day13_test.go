package day13

import "testing"

func TestDedube(t *testing.T) {
	coords := Dedupe([]Coord{Coord{1, 1}, Coord{1, 1}})
	if len(coords) != 1 {
		t.Errorf("Dedupe incorrect")
	}
}
