package main

import (
	"fmt"
	"github.com/aschwinwester/subdown/subtitles"
)

func main() {

	var options subtitles.Options
	options.Parse()
	fmt.Println(options.Title)

	tvshows, _ := subtitles.FindStreamOpenSubtitles(options.Title, options.Kind, options.Year)
	fmt.Println(len(tvshows))
	fmt.Println(tvshows[0].Title)	

}