package gozart

import (
	"fmt"
)

type Scale struct {
	name string
	key Note
	intervals []int
	Notes []Note
	//chords []Chord
}

var scales = map[string]func(*Note) *Scale{
	"chromatic": chromaticScale,
	"major": majorScale,
}

func majorScale(key *Note) *Scale {
	var direction int
	var nextNaturalNote Note
	name := "major"
	intervals := []int{2, 2, 1, 2, 2, 2, 1}
	notes := fillScale(key, intervals)

	if key.accidental == 1 || key.fullName == "F" {
		// Flat
		direction = -1
	} else {
		// Sharp
		direction = 0
	}

	// Rework notes so they wouldn't repeat
	previous, _ := NewNote(key.name)
	for i := 1; i < len(notes); i++ {
		nextNaturalNote = previous.Higher(1)
		for i := 0; i < 4; i++ {
			if nextNaturalNote.accidental != 0 || previous.name == nextNaturalNote.name {
				nextNaturalNote = nextNaturalNote.Higher(1)
			}
		}

		if nextNaturalNote.fullName != notes[i].name {
			notes[i].switchAccidental(direction, nextNaturalNote.fullName)
		}

		previous, _ = NewNote(nextNaturalNote.name)
	}

	return &Scale{
		name: name,
		key: *key,
		intervals: intervals,
		Notes: notes,
	}
}

func chromaticScale(key *Note) *Scale {
	name := "chromatic"
	intervals := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	notes := fillScale(key, intervals)

	return &Scale{
		name: name,
		key: *key,
		intervals: intervals,
		Notes: notes,
	}
}

func fillScale(key *Note, intervals []int) []Note {
	notes := make([]Note, len(intervals)+1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i+1] = notes[i].Higher(interval)
	}
	return notes
}

func NewScale(name string, key *Note) (*Scale, error) {
	if scale, ok := scales[name]; ok {
		return scale(key), nil
	} else {
		return nil, fmt.Errorf("Scale %s is not found", name)
	}
}

