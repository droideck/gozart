package gozart

import "testing"

var notes = map[string]Note{
	"A":  {FullName: "A", name: "A", accidental: 0, number: 9},
	"A#": {FullName: "A#", name: "A", accidental: 2, number: 10},
	"A♯": {FullName: "A#", name: "A", accidental: 2, number: 10},
	"Ab": {FullName: "Ab", name: "A", accidental: 1, number: 8},
	"A♭": {FullName: "Ab", name: "A", accidental: 1, number: 8},
}

var noteA, _ = NewNote("A")
var notesASharp = map[int]string{0: "A", 2: "B", 3: "C", 7: "E", 12: "A", 13: "A#", 20: "F"}
var notesAFlat = map[int]string{0: "A", 3: "Gb", 9: "C", 10: "B", 12: "A", 13: "Ab", 20: "Db"}

func TestNewNote(t *testing.T) {
	for noteName, note := range notes {
		testNote, err := NewNote(noteName)
		if err != nil {
			t.Errorf("Key search. %s", err)
		}
		if *testNote != note {
			t.Errorf("New note %v is not %v", testNote, note)
		}
	}
}

func TestNoteHigher(t *testing.T) {
	for i, shiftedNote := range notesASharp {
		higherNote := noteA.Higher(i)
		expectedNote, _ := NewNote(shiftedNote)
		if higherNote != *expectedNote {
			t.Errorf("Shifted higher note %v is not %v", higherNote, *expectedNote)
		}
	}
}

func TestNoteLower(t *testing.T) {
	for i, shiftedNote := range notesAFlat {
		lowerNote := noteA.Lower(i)
		expectedNote, _ := NewNote(shiftedNote)
		if lowerNote != *expectedNote {
			t.Errorf("Shifted lower note %v is not %v", lowerNote, *expectedNote)
		}
	}
}
