package pocketlog_test

import (
	"go-pocket-sized-projects/logger/pocketlog"
	"os"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(
		pocketlog.LevelDebug,
		pocketlog.WithOutput(os.Stdout),
	)

	debugLogger.Debugf("Teste %s", "debug")
	// Output: Teste debug
}

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)

	return len(p), nil
}

func TestLogger(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	const (
		debugMessage = "debug"
		infoMessage  = "info"
		errorMessage = "error"
	)

	testCases := map[string]testCase{
		"debug": {
			pocketlog.LevelDebug,
			debugMessage + "\n" +
				infoMessage + "\n" +
				errorMessage + "\n",
		},
		"info": {
			pocketlog.LevelInfo,
			infoMessage + "\n" +
				errorMessage + "\n",
		},
		"error": {
			pocketlog.LevelError,
			errorMessage + "\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			logger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			logger.Debugf(debugMessage)
			logger.Infof(infoMessage)
			logger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("Invalid contents. Expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
