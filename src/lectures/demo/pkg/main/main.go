package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("I'm from display")
	msg.Exciting("this is exciting")
}
