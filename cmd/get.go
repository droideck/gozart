// Copyright © 2017 Semen Pichugin <simon.pichugin@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/droideck/gozart/gozart"
	"github.com/spf13/cobra"
	"log"
)

var ChordName string
var ScaleName string
var Key string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a scale or a chord",
	Long: `You can use one of the following commands
to get a scale or a chord:

gozart get scale --name="major" --key="C#"
gozart get chord --name="Cb7"

You should use "#" for "sharp" and "b" for "flat".`,
	Run: nil,
}

var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Get a chord",
	Long: `Gives you chord notes:

gozart get chord --name="Cb7"

You should use "#" for "sharp" and "b" for "flat".`,
	Run: func(cmd *cobra.Command, args []string) {
		var notes string
		var noteName string
		var quality string

		// Parse the chord name
		if len(ChordName) > 1 {
			if string(ChordName[1]) == "#" || string(ChordName[1]) == "b" {
				noteName = ChordName[:2]
				quality = ChordName[2:]
			} else {
				noteName = ChordName[:1]
				quality = ChordName[1:]
			}

		} else {
			noteName = ChordName
			quality = ""
		}

		note, err := gozart.NewNote(noteName)
		if err != nil {
			log.Fatalf("Note search. %s", err)
		}

		chord, err := gozart.NewChord(*note, quality)
		if err != nil {
			log.Fatal(err)
		}

		for _, note := range chord.Notes {
			notes += note.FullName + " "
		}
		fmt.Println(notes)
	},
}

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Get a scale",
	Long: `Gets you a scale based on the name and key.
You can specify the next scale names:
- major
- natural minor
- melodic minor
- harmonic minor
- major pentatonic
- minor pentatonic
- blues
- dorian
- phrygian
- lydian
- mixolydian
- locryan

Use major for ionian
Use minor for aeolian

Example:
gozart get scale --name="major" --key="C#"

You should use "#" for "sharp" and "b" for "flat".`,
	Run: func(cmd *cobra.Command, args []string) {
		var notes string

		keyNote, err := gozart.NewNote(Key)
		if err != nil {
			log.Fatalf("Key search. %s", err)
		}

		scale, err := gozart.NewScale(ScaleName, keyNote)
		if err != nil {
			log.Fatal(err)
		}

		for _, note := range scale.Notes {
			notes += note.FullName + " "
		}
		fmt.Println(notes)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.AddCommand(scaleCmd)
	getCmd.AddCommand(chordCmd)
	chordCmd.PersistentFlags().StringVarP(&ChordName, "name", "n", "C7", "Chord name")
	scaleCmd.PersistentFlags().StringVarP(&ScaleName, "name", "n", "major", "Scale name")
	scaleCmd.PersistentFlags().StringVarP(&Key, "key", "k", "C", "Key name")
}
