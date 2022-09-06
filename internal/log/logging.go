package log

import "fmt"

func Error(s string) {
	Logger.Error(s)
}

func Errorf(format string, i ...interface{}) {
	Logger.Error(fmt.Sprintf(format, i...))
}
