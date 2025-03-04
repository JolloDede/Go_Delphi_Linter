package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello world")

	// r := src.NewReader()

	fi, err := os.Open("Main.pas")

	if err != nil {
		return
	}

	b := make([]byte, 20)
	for {
		n, err := fi.Read(b)

		if err != nil {
			return
		}

		fmt.Println(string(b[:n]))
	}
}
