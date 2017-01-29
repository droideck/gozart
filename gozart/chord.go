package gozart

import "fmt"

type Chord struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
}

var chordQualities = map[string][]int{
	"":       {4, 7},     // Major
	"m":      {3, 7},     // Minor
	"dim":    {3, 6},     // Diminished
	"aug":    {4, 8},     // Augmented
	"sus2":   {2, 7},     // Suspended 2nd
	"sus4":   {5, 7},     // Suspended 4th
	"M6":     {4, 7, 9},  // Minor 6th
	"m6":     {3, 7, 9},  // Minor 6th
	"7":      {4, 7, 10}, // Dominant 7th
	"M7":     {4, 7, 11}, // Major 7th
	"m7":     {3, 7, 10}, // Minor 7th
	"mM7":    {3, 7, 11}, // Minor-major 7th
	"dim7":   {3, 6, 9},  // Diminished 7th
	"hdim7":  {3, 6, 10}, // Half-diminished 7th
	"aug7":   {4, 8, 10}, // Augmented 7th
	"augM7":  {4, 8, 11}, // Augmented-major 7th
	"7sus4":  {5, 7, 10}, // Dominant 7th Suspended 4th
	"M7sus2": {2, 7, 11}, // Major 7th Suspended 2nd
	"M7sus4": {5, 7, 11}, // Major 7th Suspended 4th
}

func NewChord(key Note, quality string) (*Chord, error) {
	if _, ok := chordQualities[quality]; !ok {
		return nil, fmt.Errorf("Chord quality %s is not found", Mode)
	}

	name := key.name
	name += quality

	intervals := chordQualities[quality]
	notes := make([]Note, len(intervals)+1)

	notes[0] = key
	for i, interval := range intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		name:      name,
		key:       key,
		intervals: intervals,
		Notes:     notes,
	}, nil
}
