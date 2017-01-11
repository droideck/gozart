package gozart

var noteNamesSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var noteNamesFlat = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var noteToNum = map[string]int{"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11,
	"Db": 1, "Eb": 3, "Gb": 6, "Ab": 8, "Bb": 10,
	"C#": 1, "D#": 3, "F#": 6, "G#": 8, "A#": 10}

type Note struct {
	Name   string
	number int
	//octave int // Future
}

func (n *Note) Higher(interval int) Note {
	if nextNote := n.number + interval; nextNote < 12 {
		return Note{noteNamesSharp[nextNote], nextNote}
	} else {
		return Note{noteNamesSharp[nextNote-12], nextNote}
	}
}

func (n *Note) Lower(interval int) Note {
	if nextNote := n.number - interval; nextNote > 0 {
		return Note{noteNamesFlat[nextNote], nextNote}
	} else {
		return Note{noteNamesFlat[nextNote+12], nextNote}
	}
}

func NewNote(note string) *Note {
	return &Note{Name: note, number: noteToNum[note]}
}
