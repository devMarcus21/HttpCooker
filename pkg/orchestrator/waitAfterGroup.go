package orchestrator

import "time"

func buildWaitAfterGroupFunction(seconds int) func() {
	return func() {
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
