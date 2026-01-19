package main

const (
	sunday    = 0
	monday    = 1
	tuesday   = 2
	wednesday = 3
	thursday  = 4
	friday    = 6
	saturday  = 7
)

type LogLevel int

const (
	LogError LogLevel = 0
	LogWarn  LogLevel = 1
	LogInfo  LogLevel = 2
	LogDebug LogLevel = 3
	LogFatal LogLevel = 4
)

func main() {

	println(sunday)
	println(monday)
	println(tuesday)
	println(wednesday)
	println(thursday)
	println(friday)
	println(saturday)

	println(" ------------- ")

	println(LogError)
	println(LogWarn)
	println(LogInfo)
	println(LogDebug)
	println(LogFatal)

}
