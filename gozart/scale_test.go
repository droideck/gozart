package gozart

import (
	"reflect"
	"strings"
	"testing"
)

var trueScales = map[string]Scale{
	"chromatic.A": {
		"chromatic", Note{"A", "A", 0, 9}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]Note{
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
	"major.F": {"major", Note{"F", "F", 0, 5}, []int{2, 2, 1, 2, 2, 2, 1},
		[]Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"A", "A", 0, 9}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"D", "D", 0, 2}, {"E", "E", 0, 4}, {"F", "F", 0, 5}}},
	"major.Fb": {"major", Note{"Fb", "F", 1, 4}, []int{2, 2, 1, 2, 2, 2, 1},
		[]Note{{"Fb", "F", 1, 4}, {"Gb", "G", 1, 6}, {"Ab", "A", 1, 8}, {"Bbb", "B", 3, 9}, {"Cb", "C", 1, 11}, {"Db", "D", 1, 1}, {"Eb", "E", 1, 3}, {"Fb", "F", 1, 4}}},
	"major.B#": {"major", Note{"B#", "B", 2, 0}, []int{2, 2, 1, 2, 2, 2, 1},
		[]Note{{"B#", "B", 2, 0}, {"C##", "C", 4, 2}, {"D##", "D", 4, 4}, {"E#", "E", 2, 5}, {"F##", "F", 4, 7}, {"G##", "G", 4, 9}, {"A##", "A", 4, 11}, {"B#", "B", 2, 0}}},
	"natural minor.D": {"natural minor", Note{"D", "D", 0, 2}, []int{2, 1, 2, 2, 1, 2, 2},
		[]Note{{"D", "D", 0, 2}, {"E", "E", 0, 4}, {"F", "F", 0, 5}, {"G", "G", 0, 7}, {"A", "A", 0, 9}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"D", "D", 0, 2}}},
	"harmonic minor.E#": {"harmonic minor", Note{"E#", "E", 2, 5}, []int{2, 1, 2, 2, 1, 3, 1},
		[]Note{{"E#", "E", 2, 5}, {"F##", "F", 4, 7}, {"G#", "G", 2, 8}, {"A#", "A", 2, 10}, {"B#", "B", 2, 0}, {"C#", "C", 2, 1}, {"D##", "D", 4, 4}, {"E#", "E", 2, 5}}},
	"melodic minor.Fb": {"melodic minor", Note{"Fb", "F", 1, 4}, []int{2, 1, 2, 2, 2, 2, 1},
		[]Note{{"Fb", "F", 1, 4}, {"Gb", "G", 1, 6}, {"Abb", "A", 3, 7}, {"Bbb", "B", 3, 9}, {"Cb", "C", 1, 11}, {"Db", "D", 1, 1}, {"Eb", "E", 1, 3}, {"Fb", "F", 1, 4}}},
	"blues.F": {"blues", Note{"F", "F", 0, 5}, []int{3, 2, 1, 1, 3, 2},
		[]Note{{"F", "F", 0, 5}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"B", "B", 0, 11}, {"C", "C", 0, 0}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major pentatonic.F": {"major pentatonic", Note{"F", "F", 0, 5}, []int{2, 2, 3, 2, 3},
		[]Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D", "D", 0, 2}, {"F", "F", 0, 5}}},
	"minor pentatonic.F": {"minor pentatonic", Note{"F", "F", 0, 5}, []int{3, 2, 2, 3, 2},
		[]Note{{"F", "F", 0, 5}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major.F.aeolian": {"major", Note{"F", "F", 0, 5}, []int{2, 1, 2, 2, 1, 2, 2},
		[]Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"Db", "D", 1, 1}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major.F.locryan": {"major", Note{"F", "F", 0, 5}, []int{1, 2, 2, 1, 2, 2, 2},
		[]Note{{"F", "F", 0, 5}, {"Gb", "G", 1, 6}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"Cb", "C", 1, 11}, {"Db", "D", 1, 1}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major.F.dorian": {"major", Note{"F", "F", 0, 5}, []int{2, 1, 2, 2, 2, 1, 2},
		[]Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"D", "D", 0, 2}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major.F.phrygian": {"major", Note{"F", "F", 0, 5}, []int{1, 2, 2, 2, 1, 2, 2},
		[]Note{{"F", "F", 0, 5}, {"Gb", "G", 1, 6}, {"Ab", "A", 1, 8}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"Db", "D", 1, 1}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
	"major.F.mixolydian": {"major", Note{"F", "F", 0, 5}, []int{2, 2, 1, 2, 2, 1, 2},
		[]Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"A", "A", 0, 9}, {"Bb", "B", 1, 10}, {"C", "C", 0, 0}, {"D", "D", 0, 2}, {"Eb", "E", 1, 3}, {"F", "F", 0, 5}}},
}

func TestNewScale(t *testing.T) {
	var scaleDataSlice []string

	for scaleData, scale := range trueScales {
		scaleDataSlice = strings.Split(scaleData, ".")
		if len(scaleDataSlice) == 3 {
			Mode = scaleDataSlice[2]
		}
		keyNote, _ := NewNote(scaleDataSlice[1])
		testScale, err := NewScale(scaleDataSlice[0], keyNote)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(*testScale, scale) {
			t.Errorf("New scale %v is not %v", testScale, scale)
		}
		Mode = "ionian"
	}
}
