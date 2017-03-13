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

// jamCmd represents the jam command
var jamCmd = &cobra.Command{
	Use:   "jam",
	Short: "Get the options for an improvisation or a composition",
	Long: `Gives you chords and scales that sounds good (in some circumstances)
with the scale you give.
You can specify additional options for the command ot it will take them from a config
file or from defaults.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("jam called")
	},
}

func init() {
	RootCmd.AddCommand(jamCmd)
}
