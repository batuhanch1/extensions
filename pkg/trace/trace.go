package trace

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const TraceStart = "Starting Trace"
const TraceFinish = "Finished Trace"

type log struct {
	Duration     int64
	FunctionName string
	FileName     string
	Status       string
}

type tracer struct {
	StartTime *time.Time
	FileName  string
	FuncName  string
	SpaceLen  int
	Color     string
}

var traceMap = make(map[string]*tracer)
var spaceIndex = 0
var colorIndex = 0
var colors = []string{
	"\033[1;30m%s\033[0m",
	"\033[1;31m%s\033[0m",
	"\033[1;32m%s\033[0m",
	"\033[1;33m%s\033[0m",
	"\033[1;34m%s\033[0m",
	"\033[1;35m%s\033[0m",
	"\033[1;36m%s\033[0m",
}

var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

func Enter() string {
	var (
		programCounter uintptr
		fileName       string
		ok             bool
	)

	if programCounter, fileName, _, ok = runtime.Caller(1); !ok {
		return ""
	}

	function := runtime.FuncForPC(programCounter)
	fnName := RE_stripFnPreamble.ReplaceAllString(function.Name(), "$1")

	traceId := fmt.Sprintf("%s	%s", fileName, fnName)
	startTime := time.Now()
	traceStart(fnName, fileName, spaceIndex, colors[colorIndex])

	traceMap[traceId] = &tracer{
		StartTime: &startTime,
		FileName:  fileName,
		FuncName:  fnName,
		SpaceLen:  spaceIndex,
		Color:     colors[colorIndex],
	}

	colorIndex++
	spaceIndex += 3
	return traceId
}

func Exit(traceId string) {
	var duration int64 = 0

	trace := traceMap[traceId]
	if trace == nil {
		return
	}
	end := time.Now()
	duration = end.Sub(*trace.StartTime).Milliseconds()

	traceFinish(duration, trace.FuncName, trace.FileName, trace.SpaceLen, trace.Color)
	spaceIndex -= 3
	colorIndex--
}

func traceStart(funcName string, fileName string, spaceIndex int, color string) {
	var log = log{
		FunctionName: funcName,
		FileName:     fileName,
		Status:       TraceStart,
	}

	logMessage, _ := json.Marshal(log)

	elems := []string{
		strings.Repeat(" ", spaceIndex),
		string(logMessage),
	}

	fmt.Println(fmt.Sprintf(color, strings.Join(elems, "")))
}
func traceFinish(duration int64, funcName string, fileName string, spaceIndex int, color string) {
	var log = log{
		Duration:     duration,
		FunctionName: funcName,
		FileName:     fileName,
		Status:       TraceFinish,
	}
	logMessage, _ := json.Marshal(log)

	elems := []string{
		strings.Repeat(" ", spaceIndex),
		string(logMessage),
	}

	fmt.Println(fmt.Sprintf(color, strings.Join(elems, "")))
}
