package subtitles

import (
	"net/url"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type Kind string

type Stream struct {
	Title string `json:"name"`
	Year string `json:"year"`
	Kind string `json:"kind"`
}


func formatKind(kind string) string {
	if (kind == "tv") {
		return "serie"
	} 
	return "movie"
}
func FindStreamOpenSubtitles(title string, kind string, year uint) ([]Stream, error) {
	var Url *url.URL
    Url, err := url.Parse("https://www.opensubtitles.org")
    if err != nil {
        panic("boom")
    }

    Url.Path += "libs/suggest.php"
    parameters := url.Values{}
    parameters.Add("format", "json3")
	parameters.Add("MovieName", title)
	   
    Url.RawQuery = parameters.Encode()
	fmt.Println(Url.String())
	resp, err := http.Get(Url.String())
	if (err != nil) {
		return []Stream{}, err
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		fmt.Println(err)
	}
	
	var tvshows []Stream
	err = json.Unmarshal(body, &tvshows)
	if err != nil {
		log.Fatal(err)
	}
	filterShows := make([]Stream, 0)
	for index := 0; index < len(tvshows); index++ {
		tvshow := tvshows[index]
		if formatKind(tvshow.Kind) == kind {
			filterShows = append(filterShows, tvshow)
		} 
	}
	return filterShows, err
}





