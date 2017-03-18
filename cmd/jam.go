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
	"log"
	"github.com/spf13/cobra"
	"github.com/droideck/gozart/gozart"
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
		fmt.Println(notes)

		diatonicChords, err := scale.FindDiatonicChords("")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(diatonicChords)
	},
}

func init() {
	RootCmd.AddCommand(jamCmd)
	jamCmd.PersistentFlags().StringVarP(&ScaleName, "scale", "s", "major", "Scale name")
	jamCmd.PersistentFlags().StringVarP(&Key, "key", "k", "C", "Key name")
}
