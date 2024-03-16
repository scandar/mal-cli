package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for anime and manga",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(args)
		client := http.Client{}
		req, err := http.NewRequest("GET", "https://api.myanimelist.net/v2/users/@me/animelist?fields=list_status&limit=4", nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header = http.Header{
			"Authorization": {"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImM3ZjdjM2JlN2U4M2NhNGIyZTdhY2UzZGFiODU2NWVhYzlkYjAyYjA3ZWFkZDYwNDU5MDE0MDJjN2JkZDI5NTM0NjFkMTVlOGMzZWEzYjgyIn0.eyJhdWQiOiIyY2U2YjM5OWE0ODA4YjYxNzQxZTBlMTIxZmYyNDY2MSIsImp0aSI6ImM3ZjdjM2JlN2U4M2NhNGIyZTdhY2UzZGFiODU2NWVhYzlkYjAyYjA3ZWFkZDYwNDU5MDE0MDJjN2JkZDI5NTM0NjFkMTVlOGMzZWEzYjgyIiwiaWF0IjoxNzA5NTAwODE1LCJuYmYiOjE3MDk1MDA4MTUsImV4cCI6MTcxMjE3NTYxNSwic3ViIjoiMzUwODYxNyIsInNjb3BlcyI6W119.AgNgx6i_s-CjxOL3fM8NqxzOXff0O1FHWYCNoqHjO2EHro1Jiluv5bHOyEAvsUL6JRMo0Jg4Q3WmL4qno0TrdiuJOSnag7KMIxFYW5B3G5WapsGPbEStg0UvksW8iKmbxqpST83vs584kbsreD36vxbIf5H0aRfJb0SCw_oMQXNAMyYZ9eta8kKzuicQPheDp21rdT7896U3d1kHjOEDMmpE-eMVhEq2D_KrPAKKPiOxqDbgyLbEcrg3J6WpxqiJIJm3lTrgnl1fRFckGAvx-b0xzIYXfkvvWE2gTitcoMzv-OMR0R8gLMgxjvQjtu1iiBRukDJhUnQpSRwyLPGQGA"},
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			// bodyString := string(bodyBytes)
			var data any
			json.Unmarshal(bodyBytes, &data)

			log.Println(data)
		}

		// var data interface{}
		// if err := json.Unmarshal(res, &data); err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(data)
		// http.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImM3ZjdjM2JlN2U4M2NhNGIyZTdhY2UzZGFiODU2NWVhYzlkYjAyYjA3ZWFkZDYwNDU5MDE0MDJjN2JkZDI5NTM0NjFkMTVlOGMzZWEzYjgyIn0.eyJhdWQiOiIyY2U2YjM5OWE0ODA4YjYxNzQxZTBlMTIxZmYyNDY2MSIsImp0aSI6ImM3ZjdjM2JlN2U4M2NhNGIyZTdhY2UzZGFiODU2NWVhYzlkYjAyYjA3ZWFkZDYwNDU5MDE0MDJjN2JkZDI5NTM0NjFkMTVlOGMzZWEzYjgyIiwiaWF0IjoxNzA5NTAwODE1LCJuYmYiOjE3MDk1MDA4MTUsImV4cCI6MTcxMjE3NTYxNSwic3ViIjoiMzUwODYxNyIsInNjb3BlcyI6W119.AgNgx6i_s-CjxOL3fM8NqxzOXff0O1FHWYCNoqHjO2EHro1Jiluv5bHOyEAvsUL6JRMo0Jg4Q3WmL4qno0TrdiuJOSnag7KMIxFYW5B3G5WapsGPbEStg0UvksW8iKmbxqpST83vs584kbsreD36vxbIf5H0aRfJb0SCw_oMQXNAMyYZ9eta8kKzuicQPheDp21rdT7896U3d1kHjOEDMmpE-eMVhEq2D_KrPAKKPiOxqDbgyLbEcrg3J6WpxqiJIJm3lTrgnl1fRFckGAvx-b0xzIYXfkvvWE2gTitcoMzv-OMR0R8gLMgxjvQjtu1iiBRukDJhUnQpSRwyLPGQGA")
		// res, err := http.Get("https://api.myanimelist.net/v2/users/@me/animelist?fields=list_status&limit=4")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(res)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
