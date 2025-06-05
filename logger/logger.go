package logger

import (
	"fmt"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func LogV1(msg string) {
	formattedMsg := fmt.Sprintf("%s[+]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}

func LogV2(msg string) {
	formattedMsg := fmt.Sprintf("%s[-]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}
