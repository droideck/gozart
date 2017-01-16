package gozart

import "testing"

var notes = map[string]Note{
	"C":  {name: "C", number: 0},
	"D":  {name: "D", number: 2},
	"E":  {name: "E", number: 4},
	"F":  {name: "F", number: 5},
	"G":  {name: "G", number: 7},
	"A":  {name: "A", number: 9},
	"B":  {name: "B", number: 11},
	"C#": {name: "C#", number: 1},
	"D#": {name: "D#", number: 3},
	"F#": {name: "F#", number: 6},
	"G#": {name: "G#", number: 8},
	"A#": {name: "A#", number: 10},
	"Db": {name: "Db", number: 1},
	"Eb": {name: "Eb", number: 3},
	"Gb": {name: "Gb", number: 6},
	"Ab": {name: "Ab", number: 8},
	"Bb": {name: "Bb", number: 10},
}

var noteA, _ = NewNote("A")
var notesASharp = map[int]string{0: "A", 2: "B", 3: "C", 7: "E", 12: "A", 13: "A#"}
var notesAFlat = map[int]string{0: "A", 3: "Gb", 9: "C", 10: "B", 12: "A", 13: "Ab"}

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
		if higherNote != notes[shiftedNote] {
			t.Errorf("Shifted higher note %v is not %v", higherNote, notes[shiftedNote])
		}
	}
}

func TestNoteLower(t *testing.T) {
	for i, shiftedNote := range notesAFlat {
		lowerNote := noteA.Lower(i)
		if lowerNote != notes[shiftedNote] {
			t.Errorf("Shifted higher note %v is not %v", lowerNote, notes[shiftedNote])
		}
	}
}
