package gozart

import (
	"reflect"
	"strings"
	"testing"
)

// TODO: Add tests for new chords
var trueChords = map[string]Chord{
	"F.major":          {Note{"F", "F", 0, 5}, Quality{"major", "", []int{4, 7}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}}, 1},
	"F.diminished 7th": {Note{"F", "F", 0, 5}, Quality{"diminished 7th", "dim7", []int{3, 6, 9}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}, {"D", "D", 0, 2}}, 11},
	"F.dominant 9th":   {Note{"F", "F", 0, 5}, Quality{"dominant 9th", "9", []int{4, 7, 10, 14}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}, {"G", "G", 0, 7}}, 24},
}

func TestNewChord(t *testing.T) {
	var chordDataSlice []string

	for chordData, chord := range trueChords {
		chordDataSlice = strings.Split(chordData, ".")
		if len(chordDataSlice) == 3 {
			Mode = chordDataSlice[2]
		}
		keyNote, _ := NewNote(chordDataSlice[0])
		testChord, err := NewChord(*keyNote, chordDataSlice[1])
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(*testChord, chord) {
			t.Errorf("New chord %v is not %v", testChord, chord)
		}
		Mode = "ionian"
	}
}
