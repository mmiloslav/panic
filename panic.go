package panicline

import (
	"fmt"
	"runtime"
	"strings"
)

const runtimeDir = "/runtime/"

// GetPanicLine - retrieves and returns line where panic occured
func GetPanicLine(r interface{}) string {
	return get(r, true)
}

// GetPanicLines - retrieves and returns subsequence of lines where panic occured
func GetPanicLines(r interface{}) string {
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
