package download

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	NumComics = 2811
	// NumComics = 200
	URL     = "https://xkcd.com/"
	Postfix = "/info.0.json"
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
	fmt.Println("Got comic no. ", comic)
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

func GetComic(comic int) (ComicInfo, error) {
	comicRes, err := Download(comic)
	url := URL + fmt.Sprint(comic)
	if err != nil {
		return ComicInfo{URL: url, Err: err}, err
	}
	comicInfo := ComicInfo{URL: url, JSON: *comicRes}

	return comicInfo, nil
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
	// Get the comics
	// Consider limiting the number of simultaneous goroutines here.
	var allComics []ComicInfo
	for i := 0; i < NumComics; i++ {
		comic, comicErr := GetComic(i + 1)
		if comicErr != nil {
			fmt.Println("Had trouble getting a comic: ", comicErr)
		}
		allComics = append(allComics, comic)
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
