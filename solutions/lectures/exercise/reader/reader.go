//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodbye = "bye"
)

func main() {
	//* Use any Reader implementation from the stdlib to implement the program
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0
	for scanner.Scan() {
		//* Whenever the user types a "Q" or "q", the program should exit.
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			//* When the user enters either "hello" or "bye", the program
			//  should respond with a message after pressing the enter key.
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("command response: hi")
			case CmdGoodbye:
				numCommands += 1
				fmt.Println("command response: bye")
			}
			if text != "" {
				numLines += 1
			}
		}
	}
	//* Upon program exit, some usage statistics should be printed
	//  ('Q' and 'q' do not count towards these statistics):
	//  - The number of non-blank lines entered
	//  - The number of commands entered
	fmt.Printf("You entered %v lines\n", numLines)
	fmt.Printf("You entered %v commands\n", numCommands)
}
