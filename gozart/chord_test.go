package gozart

import (
	"reflect"
	"strings"
	"testing"
)

var trueChords = map[string]Chord{
	"F.":     {"F", Note{"F", "F", 0, 5}, []int{4, 7}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}}},
	"F.m":    {"Fm", Note{"F", "F", 0, 5}, []int{3, 7}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}}},
	"F.dim":  {"Fdim", Note{"F", "F", 0, 5}, []int{3, 6}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}}},
	"F.aug":  {"Faug", Note{"F", "F", 0, 5}, []int{4, 8}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}}},
	"F.sus2": {"Fsus2", Note{"F", "F", 0, 5}, []int{2, 7}, []Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"C", "C", 0, 0}}},
	"F.sus4": {"Fsus4", Note{"F", "F", 0, 5}, []int{5, 7}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}}},
	"F.7":    {"F7", Note{"F", "F", 0, 5}, []int{4, 7, 10}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}},
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
