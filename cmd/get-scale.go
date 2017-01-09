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

var Scale string
var Key string

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
		fmt.Println("Scale is", Scale)
		fmt.Println("Key is", Key)
	},
}
func init() {
	RootCmd.AddCommand(getScaleCmd)
	RootCmd.PersistentFlags().StringVar(&Scale, "scale", "Major", "Scale name")
	RootCmd.PersistentFlags().StringVar(&Key, "key", "C", "Key name")
}
