package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	NumComics = 2811
	URL       = "https://xkcd.com/"
	Postfix   = "/info.0.json"
)

type ComicJsonInfo struct {
	Title      string `json:"safe_title"`
	Transcript string
}

type ComicInfo struct {
	URL  string
	JSON ComicJsonInfo
	Err  error
}

func Download(comic int) (*ComicJsonInfo, error) {
	resp, err := http.Get(URL + fmt.Sprint(comic) + Postfix)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var comicInfo ComicJsonInfo
	err = json.Unmarshal(body, &comicInfo)
	if err != nil {
		return nil, err
	}
	return &comicInfo, nil
}

func GetComic(comic int, ch chan<- ComicInfo) {
	comicRes, err := Download(comic)
	url := "thing"
	if err != nil {
		ch <- ComicInfo{URL: url, Err: err}
		return
	}
	comicInfo := ComicInfo{URL: url, JSON: *comicRes}

	ch <- comicInfo
}

func main() {
	testComics := [3]int{571, 572, 573}
	ch := make(chan ComicInfo)
	for _, comic := range testComics {
		go GetComic(comic, ch)
	}

	// ensure the dir "lib" exists
	if _, err := os.Stat("lib"); os.IsNotExist(err) {
		os.Mkdir("lib", 0755)
	}

	// Open a new file in append mode
	file, err := os.OpenFile("lib/comics.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}

	defer file.Close()

	// Write the titles to the file
	for range testComics {
		comic := <-ch
		if comic.Err != nil {
			fmt.Println("error fetchinc comic: ", err)
		}
		jsonData, err := json.Marshal(comic)

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
		}
	}
}
