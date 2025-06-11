package main
import "fmt"

//// Hello returns a greeting for the named person.
func Hello(name string) string {
	message:= fmt.Sprint("HI, %v. Welcome!", name)
	return message
}

