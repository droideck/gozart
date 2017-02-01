package gozart

import (
	"reflect"
	"strings"
	"testing"
)

var trueChords = map[string]Chord{
	"F.":       {"F", Note{"F", "F", 0, 5}, []int{4, 7}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}}, 1},
	"F.m":      {"Fm", Note{"F", "F", 0, 5}, []int{3, 7}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}}, 2},
	"F.sus2":   {"Fsus2", Note{"F", "F", 0, 5}, []int{2, 7}, []Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"C", "C", 0, 0}}, 3},
	"F.sus4":   {"Fsus4", Note{"F", "F", 0, 5}, []int{5, 7}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}}, 4},
	"F.dim":    {"Fdim", Note{"F", "F", 0, 5}, []int{3, 6}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}}, 5},
	"F.aug":    {"Faug", Note{"F", "F", 0, 5}, []int{4, 8}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}}, 6},
	"F.7":      {"F7", Note{"F", "F", 0, 5}, []int{4, 7, 10}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 7},
	"F.M7":     {"FM7", Note{"F", "F", 0, 5}, []int{4, 7, 11}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 8},
	"F.m7":     {"Fm7", Note{"F", "F", 0, 5}, []int{3, 7, 10}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 9},
	"F.mM7":    {"FmM7", Note{"F", "F", 0, 5}, []int{3, 7, 11}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 10},
	"F.dim7":   {"Fdim7", Note{"F", "F", 0, 5}, []int{3, 6, 9}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}, {"D", "D", 0, 2}}, 10},
	"F.hdim7":  {"Fhdim7", Note{"F", "F", 0, 5}, []int{3, 6, 10}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}, {"D#", "D", 2, 3}}, 10},
	"F.aug7":   {"Faug7", Note{"F", "F", 0, 5}, []int{4, 8, 10}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}, {"D#", "D", 2, 3}}, 10},
	"F.augM7":  {"FaugM7", Note{"F", "F", 0, 5}, []int{4, 8, 11}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}, {"E", "E", 0, 4}}, 10},
	"F.7sus4":  {"F7sus4", Note{"F", "F", 0, 5}, []int{5, 7, 10}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 10},
	"F.M7sus2": {"FM7sus2", Note{"F", "F", 0, 5}, []int{2, 7, 11}, []Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 10},
	"F.M7sus4": {"FM7sus4", Note{"F", "F", 0, 5}, []int{5, 7, 11}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 10},
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
