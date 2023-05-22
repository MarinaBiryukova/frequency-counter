package main

import (
	"flag"
	"fmt"
	"frequency_counter/internal/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "Path to input file")
	flag.Parse()
	if len(*filename) == 0 {
		log.Fatal("Path to input file is not specified")
	}
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	result, err := util.CountFrequency(file)
	if err != nil {
		log.Fatal(err)
	}
	caser := cases.Title(language.Russian)
	for key, value := range result {
		fmt.Printf("%s: %d\n", caser.String(key), value)
	}
}
