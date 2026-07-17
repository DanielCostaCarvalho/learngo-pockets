package pocketlog_test

import (
	"go-pocket-sized-projects/logger/pocketlog"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)

	debugLogger.Debugf("Teste %s", "debug")
	// Output: Teste debug
}
