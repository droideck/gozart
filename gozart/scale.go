package gozart

import (
	"fmt"
	"sort"
)

type Scale struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
}

var Mode string = "ionian"

var modeIntervals = map[string][]int{
	"ionian":     {2, 2, 1, 2, 2, 2, 1},
	"aeolian":    {2, 1, 2, 2, 1, 2, 2},
	"locryan":    {1, 2, 2, 1, 2, 2, 2},
	"dorian":     {2, 1, 2, 2, 2, 1, 2},
	"phrygian":   {1, 2, 2, 2, 1, 2, 2},
	"lydian":     {2, 2, 2, 1, 2, 2, 1},
	"mixolydian": {2, 2, 1, 2, 2, 1, 2},
}

var scales = map[string]func(*Note) *Scale{
	"chromatic":        chromaticScale,
	"major":            majorScale,
	"natural minor":    naturalMinorScale,
	"harmonic minor":   harmonicMinorScale,
	"melodic minor":    melodicMinorScale,
	"blues":            bluesScale,
	"major pentatonic": majorPentatonicScale,
	"minor pentatonic": minorPentatonicScale,
}

func (s *Scale) FindChords() []Chord {
	var chords []Chord

	for quality := range ChordQualities {
		for _, note := range s.Notes {
			if _, ok := ChordPriorities[quality]; ok {
				chord, _ := NewChord(note, quality)
				chords = append(chords, *chord)
			}
		}
	}
	sort.Sort(ByPriority(chords))
	return chords
}

func majorScale(key *Note) *Scale {
	name := "major"
	intervals := modeIntervals[Mode]
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
	name := "natural minor"
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
	name := "harmonic minor"
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

// TODO: Implement descending melodic minor scale (it has another notes)
func melodicMinorScale(key *Note) *Scale {
	name := "melodic minor"
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

func bluesScale(key *Note) *Scale {
	name := "blues"
	intervals := []int{3, 2, 1, 1, 3, 2}
	notes := make([]Note, len(intervals)+1)
	majorNotes := majorScale(key).Notes

	notes[0] = majorNotes[0]
	notes[1] = majorNotes[2].Lower(1)
	notes[2] = majorNotes[3]
	notes[3] = majorNotes[4].Lower(1)
	notes[4] = majorNotes[4]
	notes[5] = majorNotes[6].Lower(1)
	notes[6] = majorNotes[7]

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func majorPentatonicScale(key *Note) *Scale {
	name := "major pentatonic"
	intervals := []int{2, 2, 3, 2, 3}
	notes := make([]Note, len(intervals)+1)
	majorNotes := majorScale(key).Notes

	notes[0] = majorNotes[0]
	notes[1] = majorNotes[1]
	notes[2] = majorNotes[2]
	notes[3] = majorNotes[4]
	notes[4] = majorNotes[5]
	notes[5] = majorNotes[7]

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func minorPentatonicScale(key *Note) *Scale {
	name := "minor pentatonic"
	intervals := []int{3, 2, 2, 3, 2}
	notes := make([]Note, len(intervals)+1)
	majorNotes := naturalMinorScale(key).Notes

	notes[0] = majorNotes[0]
	notes[1] = majorNotes[2]
	notes[2] = majorNotes[3]
	notes[3] = majorNotes[4]
	notes[4] = majorNotes[6]
	notes[5] = majorNotes[7]

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
	notes := make([]Note, len(intervals)+1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i+1] = notes[i].Higher(interval)
	}
	return notes
}

func NewScale(name string, key *Note) (*Scale, error) {
	if _, ok := modeIntervals[Mode]; !ok {
		return nil, fmt.Errorf("Scale mode %s is not found", Mode)
	}

	if scale, ok := scales[name]; ok {
		return scale(key), nil
	} else {
		return nil, fmt.Errorf("Scale %s is not found", name)
	}
}
