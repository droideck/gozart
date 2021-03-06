package gozart

import (
	"fmt"
)

type Scale struct {
	name      string
	key       Note
	intervals []int
	Notes     []Note
}

var scales = map[string]func(*Note) *Scale{
	"chromatic":        chromaticScale,
	"major":            majorScale,
	"dorian":           dorianScale,
	"phrygian":         phrygianScale,
	"lydian":           lydianScale,
	"mixolydian":       mixolydianScale,
	"locryan":          locryanScale,
	"natural minor":    naturalMinorScale,
	"harmonic minor":   harmonicMinorScale,
	"melodic minor":    melodicMinorScale,
	"blues":            bluesScale,
	"major pentatonic": majorPentatonicScale,
	"minor pentatonic": minorPentatonicScale,
}

func (s *Scale) FindDiatonicChords(extended string) ([]Chord, error) {
	var chords []Chord
	var numIntervals int = 3

	switch extended {
	case "7":
		numIntervals += 1
	case "9":
		numIntervals += 2
	case "11":
		numIntervals += 3
	case "13":
		numIntervals += 4
	}

	for j, note := range s.Notes[:7] {
		var chordIntervals []int

		for i := 2; i < numIntervals*2; i += 2 {
			var nextChordNote Note

			// Work out the Notes range overflow
			if j+i > 7 {
				nextChordNote = s.Notes[j+i-7]
			} else {
				nextChordNote = s.Notes[j+i]
			}

			interval, err := getInterval(nextChordNote, note)
			if err != nil {
				return nil, fmt.Errorf("%s for higher note %s and lower note %s", err, s.Notes[i].FullName, note.FullName)
			}

			// Add 12 to any interval higher then an octave
			if i < 7 {
				chordIntervals = append(chordIntervals, interval)
			} else {
				chordIntervals = append(chordIntervals, interval+12)
			}
		}
		chord, _ := NewChord(note, chordIntervals)
		chords = append(chords, *chord)
	}

	return chords, nil
}

func PrintDiatonicChords(chords []Chord) {
	var minor bool = false
	var num string

	if chords[0].quality.name[:5] == "minor" {
		minor = true
	}
	for i, chord := range chords {
		schord := chord.Sprint()
		switch i {
		case 0:
			num = "I"
		case 1:
			num = "II"
		case 2:
			if minor {
				num = "bIII"
			} else {
				num = "III"
			}
		case 3:
			num = "IV"
		case 4:
			num = "V"
		case 5:
			if minor {
				num = "bVI"
			} else {
				num = "VI"
			}
		case 6:
			if minor {
				num = "bVII"
			} else {
				num = "VII"
			}
		}
		fmt.Printf("%s%s: %s   ", num, chord.quality.suffix, schord)
	}
	fmt.Println()
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

func dorianScale(key *Note) *Scale {
	name := "dorian"
	intervals := []int{2, 1, 2, 2, 2, 1, 2}
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

func phrygianScale(key *Note) *Scale {
	name := "phrygian"
	intervals := []int{1, 2, 2, 2, 1, 2, 2}
	naturalDirections := map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": -1, "A": -1, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func lydianScale(key *Note) *Scale {
	name := "lydian"
	intervals := []int{2, 2, 2, 1, 2, 2, 1}
	naturalDirections := map[string]int{"C": 0, "D": 0, "E": 0, "F": 0, "G": 0, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func mixolydianScale(key *Note) *Scale {
	name := "mixolydian"
	intervals := []int{2, 2, 1, 2, 2, 1, 2}
	naturalDirections := map[string]int{"C": -1, "D": 0, "E": 0, "F": -1, "G": 0, "A": 0, "B": 0}
	notes := fillScaleNotes(key, intervals)
	reworkScaleNotes(notes, naturalDirections)

	return &Scale{
		name:      name,
		key:       *key,
		intervals: intervals,
		Notes:     notes,
	}
}

func locryanScale(key *Note) *Scale {
	name := "locryan"
	intervals := []int{1, 2, 2, 1, 2, 2, 2}
	naturalDirections := map[string]int{"C": -1, "D": -1, "E": -1, "F": -1, "G": -1, "A": -1, "B": 0}
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
	var nextScaleNote Note
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
		nextScaleNote = previous.NextNatural()

		if nextScaleNote.FullName != notes[i].name {
			notes[i].switchAccidental(direction, nextScaleNote.FullName)
		}

		previous, _ = NewNote(nextScaleNote.name)
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
	if scale, ok := scales[name]; ok {
		return scale(key), nil
	} else {
		return nil, fmt.Errorf("Scale %s is not found", name)
	}
}
