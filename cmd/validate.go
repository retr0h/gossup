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

	"github.com/goss-org/goss/resource"
	"github.com/spf13/cobra"

	gossoutputs "github.com/goss-org/goss/outputs"
	"github.com/retr0h/gossup/internal"
	"github.com/retr0h/gossup/internal/validator"
)

func logMatcherResultGroups(results []gossoutputs.StructuredTestResult) []any {
	logGroups := make([]any, 0, len(results))
	for _, r := range results {
		if r.Result == resource.FAIL {
			m := r.MatcherResult
			group := slog.Group(r.ResourceType,
				slog.Any("Expected", m.Expected),
				slog.Any("Actual", m.Actual),
				slog.String("Message", m.Message),
				slog.String("ResourceId", r.ResourceId),
			)

			logGroups = append(logGroups, group)
		}
	}

	return logGroups
}

// validateCmd represents the validate command.
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate system",
	Long: `validate  a system's configuration
`,
	Run: func(cmd *cobra.Command, args []string) {
		var vm internal.ValidatorManager = validator.New(
			logger,
		)
		res, err := vm.Validate()
		if err != nil {
			logFatal(
				"failed to validate",
				slog.Group("",
					slog.String("err", err.Error()),
				),
			)
		}

		duration := fmt.Sprintf("%.3fs", res.Summary.TotalDuration.Seconds())
		if res.Summary.Failed > 0 {
			logger.Error(
				"summary",
				slog.Int("count", res.Summary.TestCount),
				slog.Int("failed", res.Summary.Failed),
				slog.String("duration", duration),
				slog.Group("", logMatcherResultGroups(res.Results)...),
			)

			return
		}

		logger.Info(
			"summary",
			slog.Int("count", res.Summary.TestCount),
			slog.Int("failed", res.Summary.Failed),
			slog.String("duration", duration),
		)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
