package main

import "os"

func fatal(message string) {
	println(message)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: idl2go <file.idl>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fatal(err.Error())
	}

	err = parseIDL(string(data))
	if err != nil {
		fatal(err.Error())
	}
}
