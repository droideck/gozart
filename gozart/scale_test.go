package gozart

import (
	"testing"
	"reflect"
)

var scalesA = map[string]Scale{
	"chromatic": {
		name: "chromatic",
		key:       Note{"A", 9},
		intervals: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		Notes: []Note{
			{"A", 9},
			{"A#", 10},
			{"B", 11},
			{"C", 0},
			{"C#", 1},
			{"D", 2},
			{"D#", 3},
			{"E", 4},
			{"F", 5},
			{"F#", 6},
			{"G", 7},
			{"G#", 8},
			{"A", 9},
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

		if ! reflect.DeepEqual(*testScale, scale) {
			t.Errorf("New scale %v is not %v", testScale, scale)
		}
	}
}
