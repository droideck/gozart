package gozart

import (
	"fmt"
)

type Chord struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
	priority  int
}

type ByPriority []Chord

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].priority < a[j].priority }

var ChordPriorities = map[string]int{
	"":       1,
	"m":      2,
	"sus2":   3,
	"sus4":   4,
	"dim":    5,
	"aug":    6,
	"7":      7,
	"M7":     8,
	"m7":     9,
	"M6":     10,
	"m6":     10,
	"mM7":    10,
	"dim7":   10,
	"hdim7":  10,
	"aug7":   10,
	"augM7":  10,
	"7sus4":  10,
	"M7sus2": 10,
	"M7sus4": 10,
}

var ChordQualities = map[string][]int{
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
	if _, ok := ChordQualities[quality]; !ok {
		return nil, fmt.Errorf("Chord quality %s is not found", Mode)
	}

	name := key.name
	name += quality

	intervals := ChordQualities[quality]
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
		priority:  ChordPriorities[quality],
	}, nil
}
