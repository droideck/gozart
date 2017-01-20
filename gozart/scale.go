package gozart

import (
	"fmt"
)

type Scale struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
	//chords []Chord
}

var scales = map[string]func(*Note) *Scale{
	"chromatic":         chromaticScale,
	"major":             majorScale,
	"naturalMinor":      naturalMinorScale,
	"harmonicMinor":     harmonicMinorScale,
	"melodicMinor":      melodicMinorScale,
}

func majorScale(key *Note) *Scale {
	name := "major"
	intervals := []int{2, 2, 1, 2, 2, 2, 1}
	naturalDirections := map[string]int{"C": 0, "D": 0, "E": 0, "F": -1, "G": 0, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func naturalMinorScale(key *Note) *Scale {
	name := "naturalMinor"
	intervals := []int{2, 1, 2, 2, 1, 2, 2}
	naturalDirections := map[string]int{"C": -1, "D": -1, "E": 0, "F": -1, "G": -1, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func harmonicMinorScale(key *Note) *Scale {
	name := "harmonicMinor"
	intervals := []int{2, 1, 2, 2, 1, 3, 1}
	naturalDirections := map[string]int{"C": -1, "D": -1, "E": 0, "F": -1, "G": -1, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func melodicMinorScale(key *Note) *Scale {
	name := "melodicMinor"
	intervals := []int{2, 1, 2, 2, 2, 2, 1}
	naturalDirections := map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": -1, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func chromaticScale(key *Note) *Scale {
	name := "chromatic"
	intervals := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	notes := fillScaleNotes(key, intervals)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func reworkScaleNotes(notes []Note, naturalDirections map[string]int) {
	var direction int
	var nextNaturalNote Note
	var key Note = notes[0]

	// Define the direction where we change accidentals
	if naturalDirection, ok := naturalDirections[key.fullName]; ok {
		direction = naturalDirection
	} else {
		if key.accidental%2 == 0 {
			// Sharp
			direction = 0
		} else {
			// Flat
			direction = -1
		}
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
}

func fillScaleNotes(key *Note, intervals []int) []Note {
	notes := make([]Note, len(intervals) + 1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i + 1] = notes[i].Higher(interval)
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
