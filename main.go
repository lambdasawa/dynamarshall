package main

import "os"

func main() {
	if err := Encode(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
