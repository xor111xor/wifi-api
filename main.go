package main

import "github.com/xor111xor/wifi-api/cmd"

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
