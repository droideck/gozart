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

var ScaleName string
var Key string
var ScaleMode string

var getScaleCmd = &cobra.Command{
	Use:   "get-scale",
	Short: "Provides a music scale with chords that sound good",
	Long: `With this command you can watch over some music scale
and the chords that sound good with it.

You can run it with a command:
gozart get-scale --scale=major --key=C

If you won't give it a scale or a key, it will ask for it.`,
	Run: func(cmd *cobra.Command, args []string) {
		gozart.Mode = ScaleMode

		keyNote, err := gozart.NewNote(Key)
		if err != nil {
			log.Fatalf("Key search. %s", err)
		}

		scale, err := gozart.NewScale(ScaleName, keyNote)
		if err != nil {
			log.Fatal(err)
		}

		chords := scale.FindChords()

		fmt.Println("Scale is", ScaleName)
		fmt.Println("Mode is", gozart.Mode)
		fmt.Println("Key is", Key)
		fmt.Println("Notes are", scale.Notes)
		fmt.Println("Chords are", chords)
	},
}

func init() {
	RootCmd.AddCommand(getScaleCmd)
	RootCmd.PersistentFlags().StringVar(&ScaleName, "scale", "chromatic", "Scale name")
	RootCmd.PersistentFlags().StringVar(&Key, "key", "C", "Key name")
	RootCmd.PersistentFlags().StringVar(&ScaleMode, "mode", "ionian", "Scale mode")
}
