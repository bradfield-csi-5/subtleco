package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	NumComics = 100
	URL       = "https://xkcd.com/"
	Postfix   = "/info.0.json"
	BUF_SIZE  = 20
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

func bufferedMain(NumberOfComics int, start int) {
	// Stand up a channel for goroutine
	ch := make(chan ComicInfo)

	// Get the comics
	// Consider limiting the number of simultaneous goroutines here.
	for i := 0; i < NumberOfComics; i++ {
		go GetComic(i+start+1, ch)
	}

	// ensure the dir "lib" exists
	if _, err := os.Stat("lib"); os.IsNotExist(err) {
		os.Mkdir("lib", 0755)
	}

	// Open a new file, overwrite if exists
	file, err := os.OpenFile("lib/comics.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Println("Failed to close file:", cerr)
		}
	}()

	// Get the Comics into a writeable JSON state
	allComics, err := collectComics(ch)
	if err != nil {
		fmt.Println("Failed to collect comics:", err)
		return
	}

	jsonData, err := jsonifyComics(allComics)
	if err != nil {
		fmt.Println("Failed to convert comics to JSON:", err)
		return
	}

	// Write to the file in bulk
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write to file: ", err)
	}
}

func main() {
	comicsLeft := NumComics
	for i := 0; comicsLeft > 0; i += BUF_SIZE {
	}
}
