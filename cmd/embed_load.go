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
	"fmt"
	"log/slog"
	"os"
	"path"

	// "github.com/maja42/ember"
	"github.com/spf13/cobra"
)

func getExecutableName() string {
	// squelching err
	e, _ := os.Executable()
	return path.Dir(e)
}

// embedLoadCmd represents the load command.
var (
	attachments  string
	inPath       string
	outPath      string
	embedLoadCmd = &cobra.Command{
		Use:   "load",
		Short: "load embed files",
		Long: `Load embed files in the compiled go executable
`,
		Run: func(cmd *cobra.Command, args []string) {
			// Open executable
			exe, err := os.Open(inPath)
			if err != nil {
				// need to change so defered are run
				logFatal(
					"failed to open input file",
					slog.Group("",
						slog.String("in", inPath),
						slog.String("err", err.Error()),
					),
				)
			}
			defer func() {
				fmt.Println("HERE")
				_ = exe.Close
			}()

			// Open output
			out, err := os.OpenFile(outPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o755)
			if err != nil {
				logFatal(
					"failed to open output file",
					slog.Group("",
						slog.String("out", outPath),
						slog.String("err", err.Error()),
					),
				)
			}

			defer func() {
				_ = out.Close()
				if err := recover(); err != nil { // execution failed; delete created output file
					_ = os.Remove(outPath)
				}
			}()
		},
	}
)

func init() {
	embedCmd.AddCommand(embedLoadCmd)

	embedLoadCmd.Flags().
		StringVarP(&attachments, "attachments", "a", "attachments.json", "Path to JSON file containing a list of attachments to embed")
	embedLoadCmd.MarkPersistentFlagRequired("attachments")

	embedLoadCmd.Flags().
		StringVarP(&inPath, "in", "i", getExecutableName(), "Read the target binary that should be augmented")

	embedLoadCmd.Flags().
		StringVarP(&outPath, "outPath", "o", "", "Augment the binary with attachments, and write to out with same GOOS and GOARCH as in")
	embedLoadCmd.MarkFlagRequired("outPath")
}
