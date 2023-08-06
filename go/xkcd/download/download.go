package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	NumComics = 2811
	URL       = "https://xkcd.com/"
	Postfix   = "/info.0.json"
)

type ComicInfo struct {
	Title string `json:"safe_title"`
}

func Download(comic int) (*ComicInfo, error) {
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

	var comicInfo ComicInfo
	err = json.Unmarshal(body, &comicInfo)
	if err != nil {
		return nil, err
	}
	return &comicInfo, nil
}

func GetTitle(comic int, ch chan<- string) {
	comicInfo, err := Download(comic)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	ch <- comicInfo.Title
}

func main() {
	testComics := [3]int{571, 572, 573}
	ch := make(chan string)
	for _, comic := range testComics {
		go GetTitle(comic, ch)
	}
	for range testComics {
		fmt.Println(<-ch)
	}
}
