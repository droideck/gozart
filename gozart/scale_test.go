package gozart

import (
	"reflect"
	"testing"
)

var scalesA = map[string]Scale{
	"chromatic": {
		name:      "chromatic",
		key:       Note{"A", "A", 0, 9},
		intervals: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		Notes: []Note{
			{"A", "A", 0, 9},
			{"A#", "A", 2, 10},
			{"B", "B", 0, 11},
			{"C", "C", 0, 0},
			{"C#", "C", 2, 1},
			{"D", "D", 0, 2},
			{"D#", "D", 2, 3},
			{"E", "E", 0, 4},
			{"F", "F", 0, 5},
			{"F#", "F", 2, 6},
			{"G", "G", 0, 7},
			{"G#", "G", 2, 8},
			{"A", "A", 0, 9},
		},
	},
}

func TestNewScaleA(t *testing.T) {
	keyNote, _ := NewNote("A")

	for scaleName, scale := range scalesA {
		testScale, err := NewScale(scaleName, keyNote)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(*testScale, scale) {
			t.Errorf("New scale %v is not %v", testScale, scale)
		}
	}
}
