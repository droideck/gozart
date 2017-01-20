package gozart

var naturalNotes = map[string]int{"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11}
var noteNamesSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var noteNamesFlat = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var noteToNum = map[string]int{
	"C":    0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11,
	"Cb":   11, "Db": 1, "Eb": 3, "Fb": 4, "Gb": 6, "Ab": 8, "Bb": 10,
	"Cbb":  10, "Dbb": 0, "Ebb": 2, "Fbb": 3, "Gbb": 5, "Abb": 7, "Bbb": 9,
	"Cbbb": 9, "Dbbb": 11, "Ebbb": 1, "Fbbb": 2, "Gbbb": 4, "Abbb": 6, "Bbbb": 8,
	"C#":   1, "D#": 3, "E#": 5, "F#": 6, "G#": 8, "A#": 10, "B#": 0,
	"C##":  2, "D##": 4, "E##": 6, "F##": 7, "G##": 9, "A##": 11, "B##": 1,
	"C###": 3, "D###": 5, "E###": 7, "F###": 8, "G###": 10, "A###": 0, "B###": 2,
}
