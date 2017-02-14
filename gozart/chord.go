package gozart

import (
	"fmt"
)

type Chord struct {
	key      Note
	quality  Quality
	Notes    []Note
	priority int
}

type Quality struct {
	name      string
	suffix    string
	intervals []int
}

type ByPriority []Chord

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].priority < a[j].priority }

// TODO: Implement proper note names (with flat or sharp)
var ChordQualities = map[string]Quality{
	"major":                      {"major", "", []int{4, 7}},
	"minor":                      {"minor", "m", []int{3, 7}},
	"diminished":                 {"diminished", "dim", []int{3, 6}},
	"augmented":                  {"augmented", "aug", []int{4, 8}},
	"suspended 2nd":              {"suspended 2nd", "sus2", []int{2, 7}},
	"suspended 4th":              {"suspended 4th", "sus4", []int{5, 7}},
	"dominant 7th":               {"dominant 7th", "7", []int{4, 7, 10}},
	"major 7th":                  {"major 7th", "M7", []int{4, 7, 11}},
	"minor 7th":                  {"minor 7th", "m7", []int{3, 7, 10}},
	"minor-major 7th":            {"minor-major 7th", "mM7", []int{3, 7, 11}},
	"diminished 7th":             {"diminished 7th", "dim7", []int{3, 6, 9}},
	"half-diminished 7th":        {"half-diminished 7th", "hdim7", []int{3, 6, 10}},
	"augmented 7th":              {"augmented 7th", "aug7", []int{4, 8, 10}},
	"augmented-major 7th":        {"augmented-major 7th", "augM7", []int{4, 8, 11}},
	"dominant 7th suspended 4th": {"dominant 7th suspended 4th", "7sus4", []int{5, 7, 10}},
	"major 7th Suspended 2nd":    {"major 7th Suspended 2nd", "M7sus2", []int{2, 7, 11}},
	"major 7th Suspended 4th":    {"major 7th Suspended 4th", "M7sus4", []int{5, 7, 11}},
	"major 6th":                  {"major 6th", "M6", []int{4, 7, 9}},
	"minor 6th":                  {"minor 6th", "m6", []int{3, 7, 9}},
	"dominant 9th":               {"dominant 9th", "9", []int{4, 7, 10, 14}},
	"dominant minor 9th":         {"dominant minor 9th", "7b9", []int{4, 7, 10, 13}},
	"dominant 7th sharp 9th":     {"dominant 7th sharp 9th", "7#9", []int{4, 7, 10, 15}},
	"major 9th":                  {"major 9th", "M9", []int{4, 7, 11, 14}},
	"minor 9th":                  {"minor 9th", "7b9", []int{3, 7, 10, 14}},
	"dominant 11th":              {"dominant 11th", "11", []int{4, 7, 10, 14, 17}},
	"major 11th":                 {"major 11th", "M11", []int{4, 7, 11, 14, 18}},
	"minor 11th":                 {"minor 11th", "m11", []int{3, 7, 10, 14, 17}},
}

var ChordPriorities = map[string]int{
	"major":                      1,
	"minor":                      2,
	"diminished":                 3,
	"augmented":                  4,
	"suspended 2nd":              5,
	"suspended 4th":              6,
	"dominant 7th":               7,
	"major 7th":                  8,
	"minor 7th":                  9,
	"minor-major 7th":            10,
	"diminished 7th":             11,
	"half-diminished 7th":        12,
	"augmented 7th":              13,
	"augmented-major 7th":        14,
	"dominant 7th suspended 4th": 15,
	"major 7th Suspended 2nd":    16,
	"major 7th Suspended 4th":    17,
	"major 6th":                  18,
	"minor 6th":                  19,
	"dominant 9th":               20,
	"dominant minor 9th":         21,
	"dominant 7th sharp 9th":     22,
	"major 9th":                  23,
	"minor 9th":                  24,
	"dominant 11th":              25,
	"major 11th":                 26,
	"minor 11th":                 27,
}

func NewChord(key Note, quality string) (*Chord, error) {
	myQuality, ok := ChordQualities[quality]
	if !ok {
		return nil, fmt.Errorf("Chord quality %s is not found", quality)
	}

	notes := make([]Note, len(myQuality.intervals)+1)
	notes[0] = key
	for i, interval := range myQuality.intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		quality:  myQuality,
		key:      key,
		Notes:    notes,
		priority: ChordPriorities[myQuality.name],
	}, nil
}
