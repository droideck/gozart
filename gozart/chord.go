package gozart

import "fmt"

type Chord struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
}

var chordQualities = map[string][]int{
	"dim":    {3, 6},
	"minor":    {3, 7},
	"major":     {4, 7},
	"aug":     {4, 8},
	"sus2":   {2, 7},
	"sus4":     {5, 7},
}

func NewChord(quality string, key *Note) (*Chord, error) {
	if _, ok := chordQualities[quality]; !ok {
		return nil, fmt.Errorf("Chord quality %s is not found", Mode)
	}

	name := key.fullName + quality
	intervals := chordQualities[quality]
	notes := make([]Note, len(intervals)+1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}, nil
}
