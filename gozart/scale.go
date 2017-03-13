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
	"dorian":     {2, 1, 2, 2, 2, 1, 2},
	"phrygian":   {1, 2, 2, 2, 1, 2, 2},
	"lydian":     {2, 2, 2, 1, 2, 2, 1},
	"mixolydian": {2, 2, 1, 2, 2, 1, 2},
	"aeolian":    {2, 1, 2, 2, 1, 2, 2},
	"locryan":    {1, 2, 2, 1, 2, 2, 2},
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

// TODO: Add more chords here and to chord.go if required
var diatonicChords = map[string][]string{
	"major":              {"major", "minor", "minor", "major", "major", "minor", "diminished", "major"},
	"natural minor":      {"minor", "diminished", "major", "minor", "minor", "major", "major", "minor"},
	"harmonic minor":     {"minor", "diminished", "augmented", "minor", "major", "major", "diminished", "minor"},
	"melodic minor":      {"minor", "minor", "augmented", "major", "major", "diminished", "diminished", "minor"},
	"major 7th":          {"major 7th", "minor 7th", "minor 7th", "major 7th", "dominant 7th", "minor 7th", "half-diminished 7th", "major 7th"},
	"natural minor 7th":  {"minor 7th", "half-diminished 7th", "major 7th", "minor 7th", "minor 7th", "major 7th", "dominant 7th", "minor 7th"},
	"harmonic minor 7th": {"minor-major 7th", "half-diminished 7th", "augmented-major 7th", "minor 7th", "dominant 7th", "major 7th", "diminished 7th", "minor-major 7th"},
	"melodic minor 7th":  {"minor-major 7th", "minor 7th", "augmented-major 7th", "dominant 7th", "dominant 7th", "half-diminished 7th", "half-diminished 7th", "minor-major 7th"},
	"major 9th":          {"major 9th", "minor 9th", "minor 7b9", "major 9th", "dominant 9th", "minor 9th", "minor 7b5b9", "major 9th"},
}

// You can specify chord priorities in a config file
// Or can remove displayed chords from an output. See help for details
func (s *Scale) FindAllChords() []Chord {
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

// Config file wouldn't affect the diatonic chord output
func (s *Scale) FindDiatonicChords(extended string) []Chord {
	var chords []Chord
	var scaleName string = s.name
	if extended != "" {
		scaleName += " " + extended
	}

	for noteNum, quality := range diatonicChords[scaleName] {
		if _, ok := ChordPriorities[quality]; ok {
			chord, _ := NewChord(s.Notes[noteNum], quality)
			chords = append(chords, *chord)
		}
	}
	return chords
}

func majorScale(key *Note) *Scale {
	var naturalDirections map[string]int
	name := "major"
	intervals := modeIntervals[Mode]
	switch Mode {
	case "ionian":
		naturalDirections = map[string]int{"C": 0, "D": 0, "E": 0, "F": -1, "G": 0, "A": 0, "B": 0}
	case "dorian":
		naturalDirections = map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": -1, "A": 0, "B": 0}
	case "phrygian":
		naturalDirections = map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": -1, "A": -1, "B": 0}
	case "lydian":
		naturalDirections = map[string]int{"C": 0, "D": 0, "E": 0, "F": 0, "G": 0, "A": 0, "B": 0}
	case "mixolydian":
		naturalDirections = map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": 0, "A": 0, "B": 0}
	case "aeolian":
		naturalDirections = map[string]int{"C": -1, "D": -1, "E": 0, "F": -1, "G": -1, "A": 0, "B": 0}
	case "locryan":
		naturalDirections = map[string]int{"C": -1, "D": -1, "E": -1, "F": -1, "G": -1, "A": -1, "B": 0}
	}
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
	minorNotes := naturalMinorScale(key).Notes

	notes[0] = minorNotes[0]
	notes[1] = minorNotes[2]
	notes[2] = minorNotes[3]
	notes[3] = minorNotes[4]
	notes[4] = minorNotes[6]
	notes[5] = minorNotes[7]

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
	var nextChordNote Note
	var key Note = notes[0]

	// Define the direction where we change accidentals
	if naturalDirection, ok := naturalDirections[key.FullName]; ok {
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
		nextChordNote = previous.Higher(1)
		for i := 0; i < 4; i++ {
			if nextChordNote.accidental != 0 || previous.name == nextChordNote.name {
				nextChordNote = nextChordNote.Higher(1)
			}
		}

		if nextChordNote.FullName != notes[i].name {
			notes[i].switchAccidental(direction, nextChordNote.FullName)
		}

		previous, _ = NewNote(nextChordNote.name)
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
