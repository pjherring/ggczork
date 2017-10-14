package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	idx := 0
	reader := bufio.NewReader(os.Stdin)
	for true {
		l := Lines[idx]
		l.Say()

		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		resp := l.FindResponse(text)
		if resp.Response != "" {
			fmt.Printf("\n%s\n", resp.Response)
		} else {
			idx = resp.Goto
		}
	}
}
