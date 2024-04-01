package anime_service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/scandar/mal-cli/client"
)

var urls = map[string]string{
	"anime": "/anime",
}

type AnimePicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Anime struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	MainPicture AnimePicture `json:"main_picture"`
}

type AnimeList struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
}

func SearchAnime(q string, p int) (AnimeList, error) {
	offset := p * 10
	params := map[string]string{
		"q":      q,
		"offset": fmt.Sprintf("%d", offset),
		"limit":  "10",
	}
	res, err := client.Get(urls["anime"], params)
	if err != nil {
		fmt.Println(err)
		return AnimeList{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return AnimeList{}, err
	}

	animeList := AnimeList{}
	err = json.Unmarshal(body, &animeList)
	if err != nil {
		fmt.Println(err)
		return AnimeList{}, err
	}

	return animeList, nil
}
