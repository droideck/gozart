package gozart

type Scale struct {
	name string
	key Note
	intervals []int
	Notes []Note
	//chords []Chord
}

// TODO: Improve the logic here, so it will give right note accidentals
func NewScale(name string, key *Note) *Scale {
	intervals := ScaleIntervals[name]
	notes := make([]Note, len(intervals)+1)

	notes[0] = *key
	for i, interval := range intervals {
		notes[i+1] = notes[i].Higher(interval)
	}

	return &Scale{
		name: name,
		key: *key,
		intervals: ScaleIntervals[name],
		Notes: notes,
	}
}