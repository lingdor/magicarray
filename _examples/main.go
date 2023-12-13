package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("type the command: dtos, tag to run")
		return
	}
	if args[0] == "dtos" {
		dtosCommand()
	} else if args[0] == "tag" {
		tagCommand()
	} else {
		fmt.Println("no found command:", args[0])
	}
}
