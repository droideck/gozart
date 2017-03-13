package gozart

import "fmt"

var naturalNotes = map[string]int{"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11}
var noteNamesSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var noteNamesFlat = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var noteToNum = map[string]int{
	"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11,
	"Cb": 11, "Db": 1, "Eb": 3, "Fb": 4, "Gb": 6, "Ab": 8, "Bb": 10,
	"Cbb": 10, "Dbb": 0, "Ebb": 2, "Fbb": 3, "Gbb": 5, "Abb": 7, "Bbb": 9,
	"Cbbb": 9, "Dbbb": 11, "Ebbb": 1, "Fbbb": 2, "Gbbb": 4, "Abbb": 6, "Bbbb": 8,
	"C#": 1, "D#": 3, "E#": 5, "F#": 6, "G#": 8, "A#": 10, "B#": 0,
	"C##": 2, "D##": 4, "E##": 6, "F##": 7, "G##": 9, "A##": 11, "B##": 1,
	"C###": 3, "D###": 5, "E###": 7, "F###": 8, "G###": 10, "A###": 0, "B###": 2,
}

type Note struct {
	FullName   string
	name       string
	accidental int
	number     int
}

func resolveNoteData(fullName string) (map[string]interface{}, error) {
	var note string
	var accidental int
	var accSymbol string

	if len(fullName) == 0 {
		return nil, fmt.Errorf("Please, specify note")
	} else {
		note = string(fullName[0])
	}

	if _, ok := naturalNotes[note]; !ok {
		return nil, fmt.Errorf("Note %s is not found", fullName)
	}

	// Find out the accidentals (0 - natural, 1, 3, 5 - flats, 2, 4, 6 - sharps)
	switch {
	case len(fullName) == 1:
		accidental = 0
	case len(fullName) == 2:
		switch string(fullName[1]) {
		case "b":
			accidental = 1
			accSymbol = "b"
		case "#":
			accidental = 2
			accSymbol = "#"
		default:
			return nil, fmt.Errorf("Wrong accidental - %s", fullName)
		}
	case len(fullName) == 3:
		switch string(fullName[1:3]) {
		case "bb":
			accidental = 3
			accSymbol = "bb"
		case "##":
			accidental = 4
			accSymbol = "##"
		default:
			return nil, fmt.Errorf("Wrong accidental - %s", fullName)
		}
	case len(fullName) == 4:
		switch string(fullName[1:4]) {
		case "♭":
			accidental = 1
			accSymbol = "b"
		case "♯":
			accidental = 2
			accSymbol = "#"
		case "𝄫":
			accidental = 3
			accSymbol = "b"
		case "𝄪":
			accidental = 4
			accSymbol = "#"
		case "bbb":
			accidental = 5
			accSymbol = "bbb"
		case "###":
			accidental = 6
			accSymbol = "###"
		default:
			return nil, fmt.Errorf("Wrong accidental - %s", fullName)
		}
	default:
		return nil, fmt.Errorf("Wrong note name - %s", fullName)
	}

	fullName = fmt.Sprint(string(fullName[0]), accSymbol)
	data := map[string]interface{}{"name": fullName,
		"natural":    note,
		"accidental": accidental,
		"number":     noteToNum[fullName],
	}

	return data, nil
}

func NewNote(name string) (*Note, error) {
	data, err := resolveNoteData(name)
	if err != nil {
		return nil, err
	}
	return &Note{
		FullName:   data["name"].(string),
		name:       data["natural"].(string),
		accidental: data["accidental"].(int),
		number:     data["number"].(int),
	}, nil
}

func (n *Note) Higher(interval int) Note {
	var note *Note
	var nextNote string

	if nextNoteNum := n.number + interval; nextNoteNum < 12 {
		nextNote = noteNamesSharp[nextNoteNum]
	} else {
		if nextNoteNum < 24 {
			nextNote = noteNamesSharp[nextNoteNum-12]
		} else {
			nextNote = noteNamesSharp[nextNoteNum-24]
		}
	}

	note, _ = NewNote(nextNote)
	return *note
}

func (n *Note) Lower(interval int) Note {
	var note *Note
	var nextNote string

	if nextNoteNum := n.number - interval; nextNoteNum > -1 {
		nextNote = noteNamesFlat[nextNoteNum]
	} else {
		if nextNoteNum > -13 {
			nextNote = noteNamesFlat[nextNoteNum+12]
		} else {
			nextNote = noteNamesFlat[nextNoteNum+24]
		}
	}

	note, _ = NewNote(nextNote)
	return *note
}

func (n *Note) switchAccidental(direction int, naturalNoteName string) {
	var accidental int
	if _, ok := naturalNotes[naturalNoteName]; !ok {
		fmt.Errorf("Note %s is not found", naturalNoteName)
	}

	if direction < 0 {
		// Flats
		accidental = naturalNotes[naturalNoteName] - n.number
		if accidental < 0 {
			accidental += 12
		}
		n.name = naturalNoteName
		n.FullName = naturalNoteName
		for i := 0; i < accidental; i++ {
			n.FullName += "b"
		}
	} else {
		// Sharps
		accidental = n.number - naturalNotes[naturalNoteName]
		if accidental < 0 {
			accidental += 12
		}
		n.name = naturalNoteName
		n.FullName = naturalNoteName
		for i := 0; i < accidental; i++ {
			n.FullName += "#"
		}
	}
	data, _ := resolveNoteData(n.FullName)
	n.name = data["natural"].(string)
	n.accidental = data["accidental"].(int)
	n.number = data["number"].(int)
}
