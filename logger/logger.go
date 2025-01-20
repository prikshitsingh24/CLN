package logger

import (
	"github.com/fatih/color"
)

func Info(filename string, message string) {
	color.Green("%v - %v ", filename, message)
}

func Warn(filename string, message string) {
	color.Yellow("%v - %v", filename, message)
}

func Error(filename string, message string) {
	color.Red("%v - %v", filename, message)
}
