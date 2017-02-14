package gozart

import (
	"reflect"
	"strings"
	"testing"
)

// TODO: Add tests for new chords
var trueChords = map[string]Chord{
	"F.major":                      {Note{"F", "F", 0, 5}, Quality{"major", "", []int{4, 7}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}}, 1},
	"F.minor":                      {Note{"F", "F", 0, 5}, Quality{"minor", "m", []int{3, 7}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}}, 2},
	"F.diminished":                 {Note{"F", "F", 0, 5}, Quality{"diminished", "dim", []int{3, 6}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}}, 3},
	"F.augmented":                  {Note{"F", "F", 0, 5}, Quality{"augmented", "aug", []int{4, 8}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}}, 4},
	"F.suspended 2nd":              {Note{"F", "F", 0, 5}, Quality{"suspended 2nd", "sus2", []int{2, 7}}, []Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"C", "C", 0, 0}}, 5},
	"F.suspended 4th":              {Note{"F", "F", 0, 5}, Quality{"suspended 4th", "sus4", []int{5, 7}}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}}, 6},
	"F.dominant 7th":               {Note{"F", "F", 0, 5}, Quality{"dominant 7th", "7", []int{4, 7, 10}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 7},
	"F.major 7th":                  {Note{"F", "F", 0, 5}, Quality{"major 7th", "M7", []int{4, 7, 11}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 8},
	"F.minor 7th":                  {Note{"F", "F", 0, 5}, Quality{"minor 7th", "m7", []int{3, 7, 10}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 9},
	"F.minor-major 7th":            {Note{"F", "F", 0, 5}, Quality{"minor-major 7th", "mM7", []int{3, 7, 11}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 10},
	"F.diminished 7th":             {Note{"F", "F", 0, 5}, Quality{"diminished 7th", "dim7", []int{3, 6, 9}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}, {"D", "D", 0, 2}}, 11},
	"F.half-diminished 7th":        {Note{"F", "F", 0, 5}, Quality{"half-diminished 7th", "hdim7", []int{3, 6, 10}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"B", "B", 0, 11}, {"D#", "D", 2, 3}}, 12},
	"F.augmented 7th":              {Note{"F", "F", 0, 5}, Quality{"augmented 7th", "aug7", []int{4, 8, 10}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}, {"D#", "D", 2, 3}}, 13},
	"F.augmented-major 7th":        {Note{"F", "F", 0, 5}, Quality{"augmented-major 7th", "augM7", []int{4, 8, 11}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C#", "C", 2, 1}, {"E", "E", 0, 4}}, 14},
	"F.dominant 7th suspended 4th": {Note{"F", "F", 0, 5}, Quality{"dominant 7th suspended 4th", "7sus4", []int{5, 7, 10}}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}}, 15},
	"F.major 7th Suspended 2nd":    {Note{"F", "F", 0, 5}, Quality{"major 7th Suspended 2nd", "M7sus2", []int{2, 7, 11}}, []Note{{"F", "F", 0, 5}, {"G", "G", 0, 7}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 16},
	"F.major 7th Suspended 4th":    {Note{"F", "F", 0, 5}, Quality{"major 7th Suspended 4th", "M7sus4", []int{5, 7, 11}}, []Note{{"F", "F", 0, 5}, {"A#", "A", 2, 10}, {"C", "C", 0, 0}, {"E", "E", 0, 4}}, 17},
	"F.major 6th":                  {Note{"F", "F", 0, 5}, Quality{"major 6th", "M6", []int{4, 7, 9}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D", "D", 0, 2}}, 18},
	"F.minor 6th":                  {Note{"F", "F", 0, 5}, Quality{"minor 6th", "m6", []int{3, 7, 9}}, []Note{{"F", "F", 0, 5}, {"G#", "G", 2, 8}, {"C", "C", 0, 0}, {"D", "D", 0, 2}}, 19},
	"F.dominant 9th":               {Note{"F", "F", 0, 5}, Quality{"dominant 9th", "9", []int{4, 7, 10, 14}}, []Note{{"F", "F", 0, 5}, {"A", "A", 0, 9}, {"C", "C", 0, 0}, {"D#", "D", 2, 3}, {"G", "G", 0, 7}}, 20},
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
