package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Sorry, only works on unix/linux systems")
		os.Exit(2)
	}

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("rename takes in two arguments")
		os.Exit(2)
	}
	if len(args[0]) < len(args[1]) {
		fmt.Println("Bad alias, alias is longer than original command")
		os.Exit(2)
	}

	aliasCmd := "alias " + args[1] + "=\"" + args[0] + "\"\n"
	rcPath := GetrcPath()
	WriteToFile(rcPath, aliasCmd)

}

// GetrcPath returns ~/.zshrc or ~/.bash_profile depending on the users shell
func GetrcPath() string {
	shell := os.Getenv("SHELL")
	home := os.Getenv("HOME")
	if shell[len(shell)-3:] == "zsh" {
		return home + "/.zshrc"
	}
	return home + "/.bash_profile"
}

// WriteToFile appends a string to the given filepath
func WriteToFile(path string, msg string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte(msg)); err != nil {
		log.Fatal(err)
	}
}
