package main

import "fmt"

type LogLevel int

const (
	LevelError LogLevel = 0
	LevelWarn  LogLevel = 1
	LevelTrace LogLevel = 2
	LevelInfo  LogLevel = 3
	LevelDebug LogLevel = 4
)

var levelNames = []string{"Error", "Warning", "Trace", "Info", "Debug"}

func (level LogLevel) ToString() string {
	if level > 4 || level < 0 {
		return "Unknown"
	}

	return levelNames[level]
}

func PrintLogLevel(level LogLevel) {
	fmt.Printf("level Number : %d and Name : %s \n", level, level.ToString())
}

func main() {
	PrintLogLevel(LevelError)
	PrintLogLevel(LevelWarn)
	PrintLogLevel(LevelTrace)
	PrintLogLevel(LevelInfo)
	PrintLogLevel(LevelDebug)

	PrintLogLevel(12)
}
