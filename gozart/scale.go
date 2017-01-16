package gozart

import "fmt"

type Scale struct {
	name string
	key Note
	intervals []int
	Notes []Note
	//chords []Chord
}

var scales = map[string]func(*Note) *Scale{
	"chromatic": chromaticScale,
//	"major": majorScale,
}

func chromaticScale(key *Note) *Scale {
	name := "chromatic"
	intervals := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	notes := make([]Note, len(intervals)+1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i+1] = notes[i].Higher(interval)
	}

	return &Scale{
		name: name,
		key: *key,
		intervals: intervals,
		Notes: notes,
	}
}

func NewScale(name string, key *Note) (*Scale, error) {
	if scale, ok := scales[name]; ok {
		return scale(key), nil
	} else {
		return nil, fmt.Errorf("Scale %s is not found", name)
	}
}