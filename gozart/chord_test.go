package gozart

import (
	"reflect"
	"strings"
	"testing"
)

var trueChords = map[string]Chord{
	"F.major": {"F", Note{"F", "F", 0, 5}, []int{4, 7}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}}},
	"F.minor": {"Fm", Note{"F", "F", 0, 5}, []int{3, 7}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}}},
	"F.dim": {"Fdim", Note{"F", "F", 0, 5}, []int{3, 6}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}}},
	"F.aug": {"Faug", Note{"F", "F", 0, 5}, []int{4, 8}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}}},
}

func TestNewChord(t *testing.T) {
	var chordDataSlice []string

	for chordData, chord := range trueChords {
		chordDataSlice = strings.Split(chordData, ".")
		if len(chordDataSlice) == 3 {
			Mode = chordDataSlice[2]
		}
		keyNote, _ := NewNote(chordDataSlice[0])
		testChord, err := NewChord(chordDataSlice[1], *keyNote)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(*testChord, chord) {
			t.Errorf("New chord %v is not %v", testChord, chord)
		}
		Mode = "ionian"
	}
}
