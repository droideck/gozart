package gozart

import (
	"fmt"
	"reflect"
)

type Chord struct {
	key     Note
	quality Quality
	Notes   []Note
}

type Quality struct {
	name      string
	suffix    string
	intervals []int
}

var ChordQualities = []Quality{
	{"major", "", []int{4, 7}},
	{"minor", "m", []int{3, 7}},
	{"diminished", "dim", []int{3, 6}},
	{"augmented", "aug", []int{4, 8}},
	{"suspended 2nd", "sus2", []int{2, 7}},
	{"suspended 4th", "sus4", []int{5, 7}},
	{"dominant 7th", "7", []int{4, 7, 10}},
	{"major 7th", "M7", []int{4, 7, 11}},
	{"minor 7th", "m7", []int{3, 7, 10}},
	{"minor-major 7th", "mM7", []int{3, 7, 11}},
	{"diminished 7th", "dim7", []int{3, 6, 9}},
	{"half-diminished 7th", "m7b5", []int{3, 6, 10}},
	{"augmented 7th", "aug7", []int{4, 8, 10}},
	{"augmented-major 7th", "augM7", []int{4, 8, 11}},
	{"dominant 7th sus4", "7sus4", []int{5, 7, 10}},
	{"major 7th sus2", "M7sus2", []int{2, 7, 11}},
	{"major 7th sus4", "M7sus4", []int{5, 7, 11}},
	{"major 6th", "M6", []int{4, 7, 9}},
	{"minor 6th", "m6", []int{3, 7, 9}},
	{"major 9th", "M9", []int{4, 7, 11, 14}},
	{"minor 9th", "m9", []int{3, 7, 10, 14}},
	{"minor 7b9", "m7b9", []int{3, 7, 10, 13}},
	{"minor 7b5b9", "m7b5b9", []int{3, 6, 10, 13}},
	{"dominant 9th", "9", []int{4, 7, 10, 14}},
	{"dominant minor 9th", "7b9", []int{4, 7, 10, 13}},
	{"dominant 7#9", "7#9", []int{4, 7, 10, 15}},
	{"dominant 11th", "11", []int{4, 7, 10, 14, 17}},
	{"major 11th", "M11", []int{4, 7, 11, 14, 18}},
	{"minor 11th", "m11", []int{3, 7, 10, 14, 17}},
}

func getChordQuality(value interface{}) (Quality, error) {
	switch value := value.(type) {
	case string:
		for _, quality := range ChordQualities {
			if quality.suffix == value || quality.name == value {
				return quality, nil
			}
		}
		return nil, fmt.Errorf("Chord quality is not found by string: %s", value)
	case []int:
		for _, quality := range ChordQualities {
			if reflect.DeepEqual(quality.intervals, value) {
				return quality, nil
			}
		}
		return nil, fmt.Errorf("Chord quality is not found by intervals: %s", value)
	}
}

func NewChord(key Note, quality interface{}) (*Chord, error) {
	myQuality, err := getChordQuality(quality)
	if err {
		return nil, err
	}

	notes := make([]Note, len(myQuality.intervals)+1)
	notes[0] = key
	for i, interval := range myQuality.intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		quality: myQuality,
		key:     key,
		Notes:   notes,
	}, nil
}
