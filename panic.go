package panicline

import (
	"fmt"
	"runtime"
	"strings"
)

const runtimeDir = "/runtime/"

// One - retrieves and returns line where panic occurred
func One(r interface{}) string {
	return get(r, true)
}

// All - retrieves and returns subsequence of lines where panic occurred
func All(r interface{}) string {
	return get(r, false)
}

func get(r interface{}, one bool) string {
	if r == nil {
		return ""
	}

	i := 3
	panicLine := fmt.Sprintf("%s:\n", r)
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		i++

		if strings.Contains(file, runtimeDir) {
			continue
		}

		panicLine = panicLine + fmt.Sprintf("%s:%d\n", file, line)
		if one {
			break
		}
	}

	return panicLine
}
