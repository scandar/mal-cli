package anime_service

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/scandar/mal-cli/client"
)

var urls = map[string]string{
	"anime":         "/anime",
	"userAnimeList": "/users/@me/animelist",
}

type Status string

const (
	Watching    Status = "watching"
	Completed          = "completed"
	OnHold             = "on_hold"
	Dropped            = "dropped"
	PlanToWatch        = "plan_to_watch"
	None               = ""
)

type AnimePicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Anime struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	MainPicture AnimePicture `json:"main_picture"`
}

type ListStatus struct {
	Status             Status    `json:"status"`
	Score              int       `json:"score"`
	NumWatchedEpisodes int       `json:"num_watched_episodes"`
	IsRewatching       bool      `json:"is_rewatching"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Paging struct {
	Next string `json:"next"`
}

type AnimeList struct {
	Data []struct {
		Node Anime `json:"node"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}

type UserAnimeList struct {
	Data []struct {
		Node       Anime      `json:"node"`
		ListStatus ListStatus `json:"list_status"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}

func calcOffset(p int) int {
	return p * 10
}

func SearchAnime(q string, p int) (AnimeList, error) {
	offset := calcOffset(p)
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

func GetUserAnimeList(status Status, p int) (UserAnimeList, error) {
	offset := calcOffset(p)
	params := map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}
	if status != None {
		params["status"] = string(status)
	}

	res, err := client.Get(urls["userAnimeList"], params)
	if err != nil {
		fmt.Println(err)
		return UserAnimeList{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return UserAnimeList{}, err
	}

	userAnimeList := UserAnimeList{}
	err = json.Unmarshal(body, &userAnimeList)
	if err != nil {
		fmt.Println(err)
		return UserAnimeList{}, err
	}

	return userAnimeList, nil
}
