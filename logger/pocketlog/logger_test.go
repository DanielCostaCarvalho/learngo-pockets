package pocketlog_test

import "go-pocket-sized-projects/logger/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)

	debugLogger.Debugf("Teste %s", "debug")
	// Output: Teste debug
}
