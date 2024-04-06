package manga_service

import (
	"github.com/scandar/mal-cli/internal/services"
	"github.com/scandar/mal-cli/internal/services/shared_service"
)

var urls = map[string]string{
	"manga":         "/manga",
	"userMangaList": "/users/@me/mangalist",
	"updateManga":   "/manga/%d/my_list_status",
	"deleteManga":   "/manga/%d/my_list_status",
}

func SearchManga(q string, p int) (services.List, error) {
	d := shared_service.SearchDTO{Query: q, Page: p}
	return shared_service.Search(urls["manga"], d)
}

func GetUserMangaList(status services.MangaStatus, p int) (services.UserMangaList, error) {
	l := services.UserMangaList{}
	d := shared_service.GetListDTO{Status: string(status), Page: p}
	return shared_service.GetList(l, urls["userMangaList"], d)
}

func UpdateUserMangaList(id int, s services.MangaStatus, volumes int, chapters int, score int) (services.UpdateMangaListResponse, error) {
	l := services.UpdateMangaListResponse{}
	d := shared_service.UpdateListDTO{Status: string(s), Volumes: volumes, Chapters: chapters, Score: score}
	return shared_service.UpdateList(l, urls["updateManga"], id, d)
}

func DeleteManga(id int) (bool, error) {
	return shared_service.Delete(urls["deleteManga"], id)
}
