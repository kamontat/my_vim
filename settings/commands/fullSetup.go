// Copyright © 2018 Kamontat Chantrachirathumrong <kamontat.c@hotmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package command

import (
	setup "github.com/kamontat/my_settings/settings/setups"
	util "github.com/kamontat/my_settings/settings/utils"
	"github.com/spf13/cobra"
)

// fullSetupCmd represents the fullSetup command
var fullSetupCmd = &cobra.Command{
	Use:     "full-setup",
	Aliases: []string{"fs", "f"},
	Short:   "full setup the computer [WIP 10%]",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		setupObj := setup.Initial(os, !noInternet, alwayYes)

		util.GetLogger().Debug("Full-setup", "called on "+setupObj.ToString())

		setupObj.StartAdvance()
	},
}

func init() {
	rootCmd.AddCommand(fullSetupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fullSetupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fullSetupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
