package anime_service

import (
	"github.com/scandar/mal-cli/internal/services"
	"github.com/scandar/mal-cli/internal/services/shared_service"
)

var urls = map[string]string{
	"anime":         "/anime",
	"userAnimeList": "/users/@me/animelist",
	"updateAnime":   "/anime/%d/my_list_status",
	"deleteAnime":   "/anime/%d/my_list_status",
}

func SearchAnime(q string, p int) (services.List, error) {
	d := shared_service.SearchDTO{Query: q, Page: p}
	return shared_service.Search(urls["anime"], d)
}

func GetUserAnimeList(status services.AnimeStatus, p int) (services.UserAnimeList, error) {
	l := services.UserAnimeList{}
	d := shared_service.GetListDTO{Status: string(status), Page: p}
	return shared_service.GetList(l, urls["userAnimeList"], d)
}

func UpdateUserAnimeList(id int, s services.AnimeStatus, episodes int, score int) (services.UpdateAnimeListResponse, error) {
	l := services.UpdateAnimeListResponse{}
	d := shared_service.UpdateListDTO{Status: string(s), Episodes: episodes, Score: score}
	return shared_service.UpdateList(l, urls["updateAnime"], id, d)
}

func DeleteAnime(id int) (bool, error) {
	return shared_service.Delete(urls["deleteAnime"], id)
}
