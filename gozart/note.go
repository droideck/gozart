package gozart

import "fmt"

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

	// Find out the accidentals (0 - natural, 1 - flat, 2 - sharp)
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
	case len(fullName) == 4:
		switch string(fullName[1:4]) {
		case "♭":
			accidental = 1
			accSymbol = "b"
		case "♯":
			accidental = 2
			accSymbol = "#"
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

func (n *Note) SwitchAccidental() {
	if n.accidental != 0 {
		fmt.Println(n, "qwe")
		switch n.accidental {
		case 1:
			n.fullName = noteNamesSharp[n.number]
		case 2:
			n.fullName = noteNamesFlat[n.number]
		}
		data, _ := resolveNoteData(n.fullName)
		n.name = data["natural"].(string)
		n.accidental = data["accidental"].(int)
		n.number = data["number"].(int)
	}
}
