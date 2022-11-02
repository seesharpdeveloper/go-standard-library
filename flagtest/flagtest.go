package main

import (
	"flag"
	"fmt"
)

func main() {

	archPtr := flag.String("arch", "x86", "CPU Type")
	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "x64":
		fmt.Println("Running in 64 bit mode")
	}

	fmt.Println("Process Complete")
}
