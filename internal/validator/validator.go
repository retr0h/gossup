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

package validator

import (
	"bytes"
	"encoding/json"
	"log/slog"

	"github.com/goss-org/goss"
	gossoutputs "github.com/goss-org/goss/outputs"
	gossutil "github.com/goss-org/goss/util"
)

// New factory to create a new Validator instance.
func New(
	logger *slog.Logger,
) *Validator {
	return &Validator{
		logger: logger,
	}
}

// Validate performs validation of the system.
func (v *Validator) Validate() (*gossoutputs.StructuredOutput, error) {
	var out bytes.Buffer

	opts := []gossutil.ConfigOption{
		gossutil.WithMaxConcurrency(1),
		gossutil.WithResultWriter(&out),
		gossutil.WithSpecFile("goss.yaml"),
	}

	cfg, err := gossutil.NewConfig(opts...)
	if err != nil {
		return nil, err
	}

	_, err = goss.Validate(cfg)
	if err != nil {
		return nil, err
	}

	res := &gossoutputs.StructuredOutput{}
	err = json.Unmarshal(out.Bytes(), res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
