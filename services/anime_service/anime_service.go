package anime_service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/scandar/mal-cli/client"
	"github.com/scandar/mal-cli/services"
)

var urls = map[string]string{
	"anime":         "/anime",
	"userAnimeList": "/users/@me/animelist",
	"updateAnime":   "/anime/%d/my_list_status",
	"deleteAnime":   "/anime/%d/my_list_status",
}

func calcOffset(p int) int {
	return p * 10
}

func SearchAnime(q string, p int) (services.List, error) {
	offset := calcOffset(p)
	params := map[string]string{
		"q":      q,
		"offset": fmt.Sprintf("%d", offset),
		"limit":  "10",
	}
	res, err := client.Get(urls["anime"], params)
	if err != nil {
		fmt.Println(err)
		return services.List{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.List{}, err
	}

	animeList := services.List{}
	err = json.Unmarshal(body, &animeList)
	if err != nil {
		fmt.Println(err)
		return services.List{}, err
	}

	return animeList, nil
}

func GetUserAnimeList(status services.AnimeStatus, p int) (services.UserAnimeList, error) {
	offset := calcOffset(p)
	params := map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}
	if status != services.None {
		params["status"] = string(status)
	}

	res, err := client.Get(urls["userAnimeList"], params)
	if err != nil {
		fmt.Println(err)
		return services.UserAnimeList{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.UserAnimeList{}, err
	}

	userAnimeList := services.UserAnimeList{}
	err = json.Unmarshal(body, &userAnimeList)
	if err != nil {
		fmt.Println(err)
		return services.UserAnimeList{}, err
	}

	return userAnimeList, nil
}

func UpdateUserAnimeList(id int, s services.AnimeStatus, episodes int, score int) (services.UpdateAnimeListResponse, error) {
	params := map[string]string{}
	if s != services.None {
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
		return services.UpdateAnimeListResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.UpdateAnimeListResponse{}, err
	}

	updateAnimeListResponse := services.UpdateAnimeListResponse{}
	err = json.Unmarshal(body, &updateAnimeListResponse)
	if err != nil {
		fmt.Println(err)
		return services.UpdateAnimeListResponse{}, err
	}

	return updateAnimeListResponse, nil
}

func DeleteAnime(id int) (bool, error) {
	res, err := client.Delete(fmt.Sprintf(urls["deleteAnime"], id))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return res.StatusCode == 200, nil
}
