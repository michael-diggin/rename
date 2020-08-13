package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
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

	zshPath := os.Getenv("HOME") + "/.zshrc"
	f, err := os.OpenFile(zshPath, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(aliasCmd)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	//cmd := exec.Command("zsh", "-c", ". ./main.sh")

	//	err = cmd.Run()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}
