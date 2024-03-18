// Copyright (c) 2024 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package cmd

import (
	"log/slog"

	"github.com/maja42/ember"
	"github.com/spf13/cobra"
)

// embedListCmd represents the list command.
var embedListCmd = &cobra.Command{
	Use:   "list",
	Short: "List embed files",
	Long: `List the embed in the compiled go executable
`,
	Run: func(cmd *cobra.Command, args []string) {
		attachments, err := ember.Open()
		if err != nil {
			logFatal(
				"failed to read file",
				slog.Group("",
					slog.String("attachments", "attachments.json"),
					slog.String("err", err.Error()),
				),
			)
		}
		defer attachments.Close()

		logger.Info(
			"executable contains",
			slog.Int("attachments", attachments.Count()),
			// should probably fix this
			slog.Any("contents", attachments.List()),
		)
	},
}

func init() {
	embedCmd.AddCommand(embedListCmd)
}
