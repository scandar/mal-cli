package shared_service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/scandar/mal-cli/internal/client"
	"github.com/scandar/mal-cli/internal/services"
	"github.com/scandar/mal-cli/internal/utils"
)

type SearchDTO struct {
	Query string
	Page  int
}

type GetListDTO struct {
	Status string
	Page   int
}

type UpdateListDTO struct {
	Status   string
	Episodes int
	Score    int
	Volumes  int
	Chapters int
}

func Search(url string, d SearchDTO) (services.List, error) {
	offset := utils.CalcOffset(d.Page)
	params := map[string]string{
		"q":      d.Query,
		"offset": fmt.Sprintf("%d", offset),
		"limit":  "10",
	}
	res, err := client.Get(url, params)
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

	resList := services.List{}
	err = json.Unmarshal(body, &resList)
	if err != nil {
		fmt.Println(err)
		return services.List{}, err
	}

	return resList, nil
}

func GetList[L services.UserAnimeList | services.UserMangaList](l L, url string, d GetListDTO) (L, error) {
	offset := utils.CalcOffset(d.Page)
	params := map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}
	if d.Status != "" {
		params["status"] = d.Status
	}

	res, err := client.Get(url, params)
	if err != nil {
		fmt.Println(err)
		return l, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return l, err
	}

	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println(err)
		return l, err
	}

	return l, nil
}

func UpdateList[L services.UpdateAnimeListResponse | services.UpdateMangaListResponse](l L, url string, id int, d UpdateListDTO) (L, error) {
	params := map[string]string{}
	if d.Status != "" {
		params["status"] = d.Status
	}
	if d.Episodes != 0 {
		params["num_watched_episodes"] = fmt.Sprintf("%d", d.Episodes)
	}
	if d.Score != 0 {
		params["score"] = fmt.Sprintf("%d", d.Score)
	}
	if d.Chapters != 0 {
		params["num_chapters_read"] = fmt.Sprintf("%d", d.Chapters)
	}
	if d.Volumes != 0 {
		params["num_volumes_read"] = fmt.Sprintf("%d", d.Volumes)
	}

	res, err := client.Patch(fmt.Sprintf(url, id), params)
	if err != nil {
		fmt.Println(err)
		return l, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return l, err
	}

	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println(err)
		return l, err
	}

	return l, nil
}

func Delete(url string, id int) (bool, error) {
	res, err := client.Delete(fmt.Sprintf(url, id))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return res.StatusCode == 200, nil
}

func GetDetails(url string, id int) (services.Details, error) {
	params := map[string]string{
		"fields": "id,title,synopsis,mean,rank,status,genres,my_list_status,num_episodes,num_volumes,num_chapters",
	}
	res, err := client.Get(fmt.Sprintf(url, id), params)
	if err != nil {
		fmt.Println(err)
		return services.Details{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.Details{}, err
	}

	details := services.Details{}
	err = json.Unmarshal(body, &details)
	if err != nil {
		fmt.Println(err)
		return services.Details{}, err
	}

	return details, nil
}
