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

var jamCmd = &cobra.Command{
	Use:   "jam",
	Short: "Get the options for an improvisation or a composition",
	Long: `Gives you chords and scales that sounds good (in some circumstances)
with the scale you give.

By default it will show you diatonic chords without any extentions.`,
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

		diatonicChords, err := scale.FindDiatonicChords("")
		if err != nil {
			log.Fatal(err)
		}
		diatonicChords7, err := scale.FindDiatonicChords("7")
		if err != nil {
			log.Fatal(err)
		}
		diatonicChords9, err := scale.FindDiatonicChords("9")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Scale:", ScaleName, Key)
		fmt.Println("Scale notes:", notes)
		fmt.Println("Diatonic triads")
		gozart.PrintDiatonicChords(diatonicChords)
		fmt.Println("Diatonic seventh chords")
		gozart.PrintDiatonicChords(diatonicChords7)
		fmt.Println("Diatonic ninth chords")
		gozart.PrintDiatonicChords(diatonicChords9)
	},
}

func init() {
	RootCmd.AddCommand(jamCmd)
	jamCmd.PersistentFlags().StringVarP(&ScaleName, "scale", "s", "major", "Scale name")
	jamCmd.PersistentFlags().StringVarP(&Key, "key", "k", "C", "Key name")
}
