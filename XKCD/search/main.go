package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var comics []Comic
	tempId := os.Args[1]
	num, _ := strconv.Atoi(tempId)
	f, err := os.Open("index.json")
	if err != nil {
		log.Fatal("Cant open the index file", err)
	}
	json.NewDecoder(f).Decode(&comics)
	// fmt.Println(comics)
	for _, val := range comics {
		if val.Num == num {
			fmt.Println(val)
		}
	}
}

type Comic struct {
	Url        string
	Transcript string
	Num        int
}
