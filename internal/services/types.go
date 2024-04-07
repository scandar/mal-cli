package services

import "time"

type AnimeStatus string
type MangaStatus string

const (
	Watching    AnimeStatus = "watching"
	Completed               = "completed"
	OnHold                  = "on_hold"
	Dropped                 = "dropped"
	PlanToWatch             = "plan_to_watch"
	None                    = ""
)

const (
	Reading        MangaStatus = "reading"
	CompletedManga             = "completed"
	OnHoldManga                = "on_hold"
	DroppedManga               = "dropped"
	PlanToRead                 = "plan_to_read"
	NoneManga                  = ""
)

type Picture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Node struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	MainPicture Picture `json:"main_picture"`
}

type AnimeListStatus struct {
	Status             AnimeStatus `json:"status"`
	Score              int         `json:"score"`
	NumWatchedEpisodes int         `json:"num_watched_episodes"`
	IsRewatching       bool        `json:"is_rewatching"`
	UpdatedAt          time.Time   `json:"updated_at"`
}

type MangaListStatus struct {
	Status          MangaStatus `json:"status"`
	IsRereading     bool        `json:"is_rereading"`
	NumVolumesRead  int         `json:"num_volumes_read"`
	NumChaptersRead int         `json:"num_chapters_read"`
	Score           int         `json:"score"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type Paging struct {
	Next string `json:"next"`
}

type List struct {
	Data []struct {
		Node Node `json:"node"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}

type UserAnimeList struct {
	Data []struct {
		Node       Node            `json:"node"`
		ListStatus AnimeListStatus `json:"list_status"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}

type UserMangaList struct {
	Data []struct {
		Node       Node            `json:"node"`
		ListStatus MangaListStatus `json:"list_status"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}

type UpdateAnimeListResponse struct {
	Status             AnimeStatus   `json:"status"`
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

type UpdateMangaListResponse struct {
	Status          MangaStatus   `json:"status"`
	IsRereading     bool          `json:"is_rereading"`
	NumVolumesRead  int           `json:"num_volumes_read"`
	NumChaptersRead int           `json:"num_chapters_read"`
	Score           int           `json:"score"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Priority        int           `json:"priority"`
	NumTimesReread  int           `json:"num_times_reread"`
	RereadValue     int           `json:"reread_value"`
	Tags            []interface{} `json:"tags"`
	Comments        string        `json:"comments"`
}

type Details struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	MainPicture struct {
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"main_picture"`
	Synopsis string  `json:"synopsis"`
	Mean     float64 `json:"mean"`
	Rank     int     `json:"rank"`
	Status   string  `json:"status"`
	Genres   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	MyListStatus struct {
		Status             string    `json:"status"`
		Score              int       `json:"score"`
		NumEpisodesWatched int       `json:"num_episodes_watched"`
		IsRereading        bool      `json:"is_rereading"`
		NumVolumesRead     int       `json:"num_volumes_read"`
		NumChaptersRead    int       `json:"num_chapters_read"`
		IsRewatching       bool      `json:"is_rewatching"`
		UpdatedAt          time.Time `json:"updated_at"`
	} `json:"my_list_status"`
	NumEpisodes int `json:"num_episodes"`
	NumVolumes  int `json:"num_volumes"`
	NumChapters int `json:"num_chapters"`
}
