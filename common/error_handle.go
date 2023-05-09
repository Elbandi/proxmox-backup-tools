package common

import (
	"fmt"
	"io"
	"log"
)

func CheckErr(err error, str string, v ...interface{}) {
	if err != nil {
		log.Fatalf("Error %s: %s", fmt.Sprintf(str, v...), err.Error())
	}
}

func PrintErr(err error, str string, v ...interface{}) bool {
	if err != nil {
		log.Printf("Error %s: %s", fmt.Sprintf(str, v...), err.Error())
		return true
	}
	return false
}

func DeferClose(closer io.Closer, str string, v ...interface{}) bool {
	return PrintErr(closer.Close(), str, v...)
}
