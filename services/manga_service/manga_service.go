package manga_service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/scandar/mal-cli/client"
	"github.com/scandar/mal-cli/services"
)

var urls = map[string]string{
	"manga":         "/manga",
	"userMangaList": "/users/@me/mangalist",
	"updateManga":   "/manga/%d/my_list_status",
	"deleteManga":   "/manga/%d/my_list_status",
}

func calcOffset(p int) int {
	return p * 10
}

func SearchManga(q string, p int) (services.List, error) {
	offset := calcOffset(p)
	params := map[string]string{
		"q":      q,
		"offset": fmt.Sprintf("%d", offset),
		"limit":  "10",
	}
	res, err := client.Get(urls["manga"], params)
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

	mangaList := services.List{}
	err = json.Unmarshal(body, &mangaList)
	if err != nil {
		fmt.Println(err)
		return services.List{}, err
	}

	return mangaList, nil
}

func GetUserMangaList(status services.MangaStatus, p int) (services.UserMangaList, error) {
	offset := calcOffset(p)
	params := map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}
	if status != services.NoneManga {
		params["status"] = string(status)
	}

	res, err := client.Get(urls["userMangaList"], params)
	if err != nil {
		fmt.Println(err)
		return services.UserMangaList{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.UserMangaList{}, err
	}

	userMangaList := services.UserMangaList{}
	err = json.Unmarshal(body, &userMangaList)
	if err != nil {
		fmt.Println(err)
		return services.UserMangaList{}, err
	}

	return userMangaList, nil
}

func UpdateUserMangaList(id int, s services.MangaStatus, volumes int, chapters int, score int) (services.UpdateMangaListResponse, error) {
	params := map[string]string{}
	if s != services.NoneManga {
		params["status"] = string(s)
	}
	if chapters != 0 {
		params["num_chapters_read"] = fmt.Sprintf("%d", chapters)
	}
	if volumes != 0 {
		params["num_volumes_read"] = fmt.Sprintf("%d", volumes)
	}
	if score != 0 {
		params["score"] = fmt.Sprintf("%d", score)
	}

	res, err := client.Patch(fmt.Sprintf(urls["updateManga"], id), params)
	if err != nil {
		fmt.Println(err)
		return services.UpdateMangaListResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return services.UpdateMangaListResponse{}, err
	}

	updateMangaListResponse := services.UpdateMangaListResponse{}
	err = json.Unmarshal(body, &updateMangaListResponse)
	if err != nil {
		fmt.Println(err)
		return services.UpdateMangaListResponse{}, err
	}

	return updateMangaListResponse, nil
}

func DeleteManga(id int) (bool, error) {
	res, err := client.Delete(fmt.Sprintf(urls["deleteManga"], id))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return res.StatusCode == 200, nil
}
