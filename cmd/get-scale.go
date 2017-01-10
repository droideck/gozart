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
	"github.com/spf13/cobra"
)

var ScaleName string
var Key string

func getNextNote(currentNote string, step int) string {
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	noteToNum := map[string]int{"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11}
	if nextNote := noteToNum[currentNote] + step; nextNote < 12 {
		return notes[nextNote]
	} else {
		return notes[nextNote - 12]
	}
}

// TODO: Move it to distinct module
func getScale(root string, intervals []int) []string {
	scale := make([]string, len(intervals) + 1)

	currNote := root
	scale[0] = currNote
	for i, interval := range intervals {
		currNote = getNextNote(currNote, interval)
		scale[i + 1] = currNote
	}
	return scale
}

// get-scaleCmd represents the get-scale command
var getScaleCmd = &cobra.Command{
	Use:   "get-scale",
	Short: "Provides a music scale with chords that sound good",
	Long: `With this command you can watch over some music scale
and the chords that sound good with it.

You can run it with a command:
gozart get-scale --scale=major --key=C

If you won't give it a scale or a key, it will ask for it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get-scale was called")
		fmt.Println("Scale is", ScaleName)
		fmt.Println("Key is", Key)
		scaleIntervals := []int{2, 2, 1, 2, 2, 2}
		fmt.Println(getScale(Key, scaleIntervals))
	},
}

func init() {
	RootCmd.AddCommand(getScaleCmd)
	RootCmd.PersistentFlags().StringVar(&ScaleName, "scale", "Major", "Scale name")
	RootCmd.PersistentFlags().StringVar(&Key, "key", "C", "Key name")
}
