package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var comics []Comic

	f, err := os.Create("index.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// reader := gob.NewDecoder(f)
	var n bytes.Buffer
	writer := gob.NewEncoder(&n)
	for i := 1; i <= 600; i++ {
		if i == 404 {
			continue
		} else {
			comic, err := getComic(i)
			if err != nil {
				log.Fatal(err)
			}
			comics = append(comics, comic)

			// f.WriteString(fmt.Sprintf("%d", i))
			// f.WriteString("\n")
			// f.WriteString(comic.Url)
			// f.WriteString(comic.Transcript)
			// f.WriteString("\n------------------------------------------------\n")
		}
	}
	err = writer.Encode(comics)
	if err != nil {
		fmt.Println("Cant encode")
	}
	file, err := json.MarshalIndent(comics, "", " ")
	os.WriteFile("index.json", file, 0644)

}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("cant get comic from xkcd %d %s", n, resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

type Comic struct {
	Url        string
	Transcript string
	Num        int
}
