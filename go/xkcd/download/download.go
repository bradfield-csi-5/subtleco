package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// NumComics = 2811
	NumComics = 3
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
	url := URL + fmt.Sprint(comic)
	if err != nil {
		ch <- ComicInfo{URL: url, Err: err}
		return
	}
	comicInfo := ComicInfo{URL: url, JSON: *comicRes}

	ch <- comicInfo
}

func collectComics(ch chan ComicInfo) ([]ComicInfo, error) {
	var comics []ComicInfo
	for i := 0; i < NumComics; i++ {
		comic := <-ch

		if comic.Err != nil {
			fmt.Println("error fetchinc comic: ", comic.Err)
			continue
		}
		comics = append(comics, comic)
	}
	return comics, nil
}

func jsonifyComics(allComics []ComicInfo) ([]byte, error) {
	jsonData, jsonErr := json.MarshalIndent(allComics, "", "    ")
	if jsonErr != nil {
		fmt.Println("error writing info: ", jsonErr)
		return make([]byte, 0), nil
	}
	return jsonData, nil
}

func main() {
	// Stand up a channel for goroutine
	ch := make(chan ComicInfo)

	// Get the comics
	for i := 0; i < NumComics; i++ {
		go GetComic(i+1, ch)
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

	// Get the Comics into a writeable JSON state
	allComics, err := collectComics(ch)
	jsonData, err := jsonifyComics(allComics)

	// Write to the file in bulk
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write to file: ", err)
	}
}
