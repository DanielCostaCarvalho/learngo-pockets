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
	// Output: [DEBUG]: Teste debug
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
			"[DEBUG]: " + debugMessage + "\n" +
				"[INFO]: " + infoMessage + "\n" +
				"[ERROR]: " + errorMessage + "\n",
		},
		"info": {
			pocketlog.LevelInfo,
			"[INFO]: " + infoMessage + "\n" +
				"[ERROR]: " + errorMessage + "\n",
		},
		"error": {
			pocketlog.LevelError,
			"[ERROR]: " + errorMessage + "\n",
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

func TestLoggerLogf(t *testing.T) {
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
			"[DEBUG]: " + debugMessage + "\n" +
				"[INFO]: " + infoMessage + "\n" +
				"[ERROR]: " + errorMessage + "\n",
		},
		"info": {
			pocketlog.LevelInfo,
			"[INFO]: " + infoMessage + "\n" +
				"[ERROR]: " + errorMessage + "\n",
		},
		"error": {
			pocketlog.LevelError,
			"[ERROR]: " + errorMessage + "\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			logger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			logger.Logf(pocketlog.LevelDebug, debugMessage)
			logger.Logf(pocketlog.LevelInfo, infoMessage)
			logger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("Invalid contents. Expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
