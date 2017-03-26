package gozart

import (
	"fmt"
	"reflect"
)

// TODO: Add print chord method
type Chord struct {
	key     Note
	quality Quality
	Notes   []Note
}

type Quality struct {
	name      string
	suffix    string
	intervals []int
}

var ChordQualities = []Quality{
	{"major", "", []int{4, 7}},
	{"minor", "mi", []int{3, 7}},
	{"diminished", "dim", []int{3, 6}},
	{"augmented", "+", []int{4, 8}},
	{"suspended fourth", "sus", []int{5, 7}},
	{"dominant seventh", "7", []int{4, 7, 10}},
	{"major seventh", "ma7", []int{4, 7, 11}},
	{"minor seventh", "mi7", []int{3, 7, 10}},
	{"minor (major seventh)", "mi(ma7)", []int{3, 7, 11}},
	{"diminished seventh", "dim7", []int{3, 6, 9}},
	{"minor seventh (flat five)", "mi7(b5)", []int{3, 6, 10}},
	{"dominant seventh augmented fifth", "+7", []int{4, 8, 10}},
	{"augmented major seventh", "+ma7", []int{4, 8, 11}},
	{"seventh suspended fourth", "7sus", []int{5, 7, 10}},
	{"major seventh suspended fourth", "ma7sus", []int{5, 7, 11}},
	{"major sixth", "6", []int{4, 7, 9}},
	{"minor sixth", "mi6", []int{3, 7, 9}},
	{"major ninth", "ma9", []int{4, 7, 11, 14}},
	{"minor ninth", "mi9", []int{3, 7, 10, 14}},
	{"minor seventh (flat nine)", "mi7b9", []int{3, 7, 10, 13}},
	{"minor seventh (flat five, flat nine)", "mi7b5b9", []int{3, 6, 10, 13}},
	{"dominant nineth", "9", []int{4, 7, 10, 14}},
	{"dominant seventh (flat nine)", "7b9", []int{4, 7, 10, 13}},
	{"dominant seventh (sharp nine)", "7#9", []int{4, 7, 10, 15}},
	{"minor (major ninth)", "mi(ma9)", []int{3, 7, 11, 14}},
	{"augmented major ninth", "+ma9", []int{4, 8, 11, 14}},
	{"major seventh (sharp nine)", "ma7#9", []int{4, 7, 11, 15}},
	{"diminished ninth", "dim9", []int{3, 6, 9, 13}},
	{"dominant eleventh", "11", []int{4, 7, 10, 14, 17}},
	{"major eleventh", "ma11", []int{4, 7, 11, 14, 18}},
	{"minor eleventh", "mi11", []int{3, 7, 10, 14, 17}},
}

func getChordQuality(value interface{}) (*Quality, error) {
	switch value := value.(type) {
	case string:
		for _, quality := range ChordQualities {
			if quality.suffix == value || quality.name == value {
				return &quality, nil
			}
		}
	case []int:
		for _, quality := range ChordQualities {
			if reflect.DeepEqual(quality.intervals, value) {
				return &quality, nil
			}
		}
	}
	return nil, fmt.Errorf("Chord quality is not found: %s", value)
}

func (ch Chord) Sprint() string {
	var notes string = ch.Notes[0].FullName

	for _, note := range ch.Notes[1:] {
		notes += " " + note.FullName
	}
	return fmt.Sprintf("%s%s [%s]", ch.key.FullName, ch.quality.suffix, notes)
}

func NewChord(key Note, quality interface{}) (*Chord, error) {
	myQuality, err := getChordQuality(quality)
	if err != nil {
		return nil, err
	}

	notes := make([]Note, len(myQuality.intervals)+1)
	notes[0] = key
	for i, interval := range myQuality.intervals {
		notes[i+1] = notes[0].Higher(interval)
	}

	return &Chord{
		quality: *myQuality,
		key:     key,
		Notes:   notes,
	}, nil
}
