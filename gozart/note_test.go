package gozart

import (
	"testing"
)

var notes = map[string]Note{
	"C":  {Name: "C", number: 0},
	"D":  {Name: "D", number: 2},
	"E":  {Name: "E", number: 4},
	"F":  {Name: "F", number: 5},
	"G":  {Name: "G", number: 7},
	"A":  {Name: "A", number: 9},
	"B":  {Name: "B", number: 11},
	"C#": {Name: "C#", number: 1},
	"D#": {Name: "D#", number: 3},
	"F#": {Name: "F#", number: 6},
	"G#": {Name: "G#", number: 8},
	"A#": {Name: "A#", number: 10},
	"Db": {Name: "Db", number: 1},
	"Eb": {Name: "Eb", number: 3},
	"Gb": {Name: "Gb", number: 6},
	"Ab": {Name: "Ab", number: 8},
	"Bb": {Name: "Bb", number: 10},
}

var noteA = NewNote("A")
var notesASharp = map[int]string{0: "A", 2: "B", 3: "C", 7: "E", 12: "A", 13: "A#"}
var notesAFlat = map[int]string{0: "A", 3: "Gb", 9: "C", 10: "B", 12: "A", 13: "Ab"}

func TestNewNote(t *testing.T) {
	for noteName, note := range notes {
		testNote := *NewNote(noteName)
		if testNote != note {
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
