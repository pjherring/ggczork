package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var msg string
	idx := 0
	reader := bufio.NewReader(os.Stdin)

	for true {
		l := Lines[idx]

		if msg != "" {
			fmt.Printf("\n%s\n", msg)
			msg = ""
		} else {
			l.Say()
		}

		fmt.Printf("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		resp := l.FindResponse(text)
		if resp.Response != "" {
			msg = resp.Response
		} else {
			idx = resp.Goto
		}
	}
}
