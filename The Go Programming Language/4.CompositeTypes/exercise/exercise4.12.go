package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	comicList := make(map[int]string, 10)
	for i := 1; i <= 2; i++ {
		url := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		comic, err := downloadComicUrl(url)
		if err != nil {
			fmt.Println(err)
		}
		comicList[i] = comic
	}
	fmt.Println(comicList)
}

func downloadComicUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()
	return string(body), nil
}

func saveComic() {

}

func readLocalComic() {

}