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
	name := "major"
	intervals := []int{2, 2, 1, 2, 2, 2, 1}
	notes := fillScale(key, intervals)

	// Major scale shouldn't have repeated notes
	mainNotes := naturalNotes

	// Remove all natural notes from check list
	for _, note := range notes {
		if note.accidental == 0 {
			delete(mainNotes, note.fullName)
		}
	}

	// Work out accidental notes
	for i := range notes {
		if notes[i].accidental != 0 {
			_, ok := mainNotes[notes[i].name]
			if ok {
				delete(mainNotes, notes[i].name)
			} else {
				notes[i].SwitchAccidental()
			}
		}
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

