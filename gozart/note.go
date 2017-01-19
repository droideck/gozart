package gozart

import (
	"fmt"
)

type Note struct {
	fullName string
	name string
	accidental int
	number int
	//octave int // Future
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

	if _, ok := naturalNotes[note]; ! ok {
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
	"natural": note,
	"accidental": accidental,
	"number": noteToNum[fullName],
	}

	return data, nil
}

func NewNote(name string) (*Note, error) {
	data, err := resolveNoteData(name)
	if err != nil {
		return nil, err
	}
	return &Note{
		fullName: data["name"].(string),
		name: data["natural"].(string),
		accidental: data["accidental"].(int),
		number: data["number"].(int),
	}, nil
}

func (n *Note) Higher(interval int) Note {
	var note *Note
	var nextNote string

	if nextNoteNum := n.number + interval; nextNoteNum < 12 {
		nextNote = noteNamesSharp[nextNoteNum]
	} else {
		nextNote = noteNamesSharp[nextNoteNum-12]
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
		nextNote = noteNamesFlat[nextNoteNum+12]
	}

	note, _ = NewNote(nextNote)
	return *note
}

func (n *Note) switchAccidental(direction int, naturalNoteName string) {
	var accidental int
	if _, ok := naturalNotes[naturalNoteName]; ! ok {
		fmt.Errorf("Note %s is not found", naturalNoteName)
	}

	if direction < 0 {
		// Flats
		accidental = naturalNotes[naturalNoteName] - n.number
		if accidental < 0 {
			accidental += 12
		}
		n.name = naturalNoteName
		n.fullName = naturalNoteName
		for i := 0; i < accidental; i ++ {
			n.fullName += "b"
		}
	} else {
		// Sharps
		accidental = n.number - naturalNotes[naturalNoteName]
		if accidental < 0 {
			accidental += 12
		}
		n.name = naturalNoteName
		n.fullName = naturalNoteName
		for i := 0; i < accidental; i ++ {
			n.fullName += "#"
		}
	}
	data, _ := resolveNoteData(n.fullName)
	n.name = data["natural"].(string)
	n.accidental = data["accidental"].(int)
	n.number = data["number"].(int)
}
