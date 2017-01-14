package gozart

type Note struct {
	name   string
	number int
	//octave int // Future
}

func NewNote(note string) *Note {
	return &Note{name: note, number: noteToNum[note]}
}

func (n *Note) Higher(interval int) Note {
	if nextNote := n.number + interval; nextNote < 12 {
		return Note{noteNamesSharp[nextNote], nextNote}
	} else {
		return Note{noteNamesSharp[nextNote-12], nextNote - 12}
	}
}

func (n *Note) Lower(interval int) Note {
	if nextNote := n.number - interval; nextNote > -1 {
		return Note{noteNamesFlat[nextNote], nextNote}
	} else {
		return Note{noteNamesFlat[nextNote+12], nextNote + 12}
	}
}
