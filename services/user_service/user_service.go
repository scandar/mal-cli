package user_service

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/scandar/mal-cli/client"
)

var urls = map[string]string{
	"userInfo": "/users/@me",
}

type UserInfo struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Location        string    `json:"location"`
	JoinedAt        time.Time `json:"joined_at"`
	AnimeStatistics struct {
		NumItemsWatching    int     `json:"num_items_watching"`
		NumItemsCompleted   int     `json:"num_items_completed"`
		NumItemsOnHold      int     `json:"num_items_on_hold"`
		NumItemsDropped     int     `json:"num_items_dropped"`
		NumItemsPlanToWatch int     `json:"num_items_plan_to_watch"`
		NumItems            int     `json:"num_items"`
		NumDaysWatched      float64 `json:"num_days_watched"`
		NumDaysWatching     float64 `json:"num_days_watching"`
		NumDaysCompleted    float64 `json:"num_days_completed"`
		NumDaysOnHold       int     `json:"num_days_on_hold"`
		NumDaysDropped      float64 `json:"num_days_dropped"`
		NumDays             float64 `json:"num_days"`
		NumEpisodes         int     `json:"num_episodes"`
		NumTimesRewatched   int     `json:"num_times_rewatched"`
		MeanScore           float64 `json:"mean_score"`
	} `json:"anime_statistics"`
}

func GetUserInfo() (UserInfo, error) {
	params := map[string]string{
		"fields": "anime_statistics",
	}
	res, err := client.Get(urls["userInfo"], params)
	if err != nil {
		fmt.Println(err)
		return UserInfo{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return UserInfo{}, err
	}

	userInfo := UserInfo{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println(err)
		return UserInfo{}, err
	}

	return userInfo, nil
}
