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

func NewChord(quality string, key string) (*Chord, error) {
	if _, ok := chordQualities[quality]; !ok {
		return nil, fmt.Errorf("Chord quality %s is not found", Mode)
	}

	keyNote, err := NewNote(key)
	if err != nil {
		return nil, fmt.Errorf("Key note name %s is not found", key)
	}

	name := keyNote.name
	if quality != "major" {
		name += quality
	}
	intervals := chordQualities[quality]
	notes := make([]Note, len(intervals)+1)

	notes[0] = *keyNote
	for i, interval := range intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		name:      name,
		key:       *keyNote,
		intervals: intervals,
		Notes:     notes,
	}, nil
}
