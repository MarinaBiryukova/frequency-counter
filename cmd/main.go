package main

import (
	"flag"
	"fmt"
	"frequency_counter/internal/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"sort"
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

	resultMap, err := util.CountFrequency(file)
	if err != nil {
		log.Fatal(err)
	}
	resultSlice := util.MapToSlice(resultMap)
	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].Value > resultSlice[j].Value
	})
	caser := cases.Title(language.Russian)
	for _, x := range resultSlice {
		fmt.Printf("%s: %d\n", caser.String(x.Key), x.Value)
	}
}
