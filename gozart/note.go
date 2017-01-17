package gozart

import "fmt"

type Note struct {
	fullName string
	note string
	accidental int
	number int
	//octave int // Future
}

func NewNote(name string) (*Note, error) {
	var note string
	var accidental int
	var accSymbol string

	if len(name) == 0 {
		return nil, fmt.Errorf("Please, specify note")
	} else {
		note = string(name[0])
	}

	if _, ok := mainNotes[note]; ok {
	} else {
		return nil, fmt.Errorf("Note %s is not found", name)
	}

	// Find out the accidentals (0 - natural, 1 - flat, 2 - sharp)
	switch {
	case len(name) == 1:
		accidental = 0
	case len(name) == 2:
		switch string(name[1]) {
		case "b":
			accidental = 1
			accSymbol = "b"
		case "#":
			accidental = 2
			accSymbol = "#"
		default:
			return nil, fmt.Errorf("Wrong accidental - %s", name)
		}
	case len(name) == 4:
		switch string(name[1:4]) {
		case "♭":
			accidental = 1
			accSymbol = "b"
		case "♯":
			accidental = 2
			accSymbol = "#"
		default:
			return nil, fmt.Errorf("Wrong accidental - %s", name)
		}
	default:
		return nil, fmt.Errorf("Wrong note name - %s", name)
	}

	name = fmt.Sprint(string(name[0]), accSymbol)
	return &Note{
		fullName: name,
		note: note,
		accidental: accidental,
		number: noteToNum[name],
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
