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
	"updateAnime":   "/anime/%d/my_list_status",
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

type UpdateAnimeListResponse struct {
	Status             string        `json:"status"`
	Score              int           `json:"score"`
	NumEpisodesWatched int           `json:"num_episodes_watched"`
	IsRewatching       bool          `json:"is_rewatching"`
	UpdatedAt          time.Time     `json:"updated_at"`
	Priority           int           `json:"priority"`
	NumTimesRewatched  int           `json:"num_times_rewatched"`
	RewatchValue       int           `json:"rewatch_value"`
	Tags               []interface{} `json:"tags"`
	Comments           string        `json:"comments"`
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

func UpdateUserAnimeList(id int, s Status, episodes int, score int) (UpdateAnimeListResponse, error) {
	params := map[string]string{}
	if s != None {
		params["status"] = string(s)
	}
	if episodes != 0 {
		params["num_watched_episodes"] = fmt.Sprintf("%d", episodes)
	}
	if score != 0 {
		params["score"] = fmt.Sprintf("%d", score)
	}

	res, err := client.Patch(fmt.Sprintf(urls["updateAnime"], id), params)
	if err != nil {
		fmt.Println(err)
		return UpdateAnimeListResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return UpdateAnimeListResponse{}, err
	}

	updateAnimeListResponse := UpdateAnimeListResponse{}
	err = json.Unmarshal(body, &updateAnimeListResponse)
	if err != nil {
		fmt.Println(err)
		return UpdateAnimeListResponse{}, err
	}

	return updateAnimeListResponse, nil
}
