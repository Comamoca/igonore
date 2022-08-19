package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const ignoreName = ".gitignore"

func main() {
	flag.Parse()
	args := flag.Args()

	if inStr(args, "help") {
		showHelp()
		os.Exit(0)
	}

	if len(args) == 0 {
		intaractive()
	} else {
		ignore := fetch(args)
		save(ignoreName, ignore)
	}
	fmt.Println("Done! 🎉")
}

func intaractive() {
	var selectLangs []string
	for {
		err := Prompt("Do you want to add the specified language")

		if err != nil {
			break
		} else {
			fmt.Println("Open the viewfinder...")
		}

		lang := finder()
		selectLangs = append(selectLangs, lang)
	}
	ignore := fetch(selectLangs)
	save(ignoreName, ignore)
}

func save(fileName string, ctn string) {
	_, err := os.Stat(fileName)
	if err == nil {
		err := Prompt("Overwrite file. Are you sure")
		if err != nil {
			fmt.Println("Exit.")
			os.Exit(1)
		}

	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	file.Write([]byte(ctn))
}

func showHelp() {
	const help = `
  igonore - .gitignore generator written in Go

  Usage: igonore
         Generate .gitignore interactively.

         igonore langage langage2...
         Generate a .gitignore file for the specified language.

         Ex.) igonore go node
  `
	fmt.Println(help)
}

func inStr(arr []string, str string) bool {

	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
