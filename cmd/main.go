package main

import (
	"flag"
	"frequency_counter/pkg/util"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "Name of input file")
	flag.Parse()
	if len(*filename) == 0 {
		log.Fatal("Name of input file is not specified")
	}
	file, err := os.Open("resources/" + *filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	result := util.CountFrequency(file)
	for key, value := range result {
		util.PrettyPrint(key, value)
	}
}
