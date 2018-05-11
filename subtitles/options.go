package subtitles

import (
	"flag"
)
type Options struct {
	Kind string
	Title string
	Year uint
	Folder string
	Verbose bool
}
var parsed = false;

func (options *Options) Parse()  {
	if (parsed) {
		return
	}
	verbosePtr := flag.Bool("v", false, "Verbose logging")
	folderNamePtr := flag.String("folder", ".", "Folder name to store subtitles")
	titlePtr := flag.String("title", ".", "Title or name of movie or tvshow")
	kindPtr := flag.String("kind", "serie", "choice between movie or serie")
	yearPtr := flag.Uint("year", 0, "Originated in year")
	
	flag.Parse()

	options.Folder = *folderNamePtr
	options.Kind = *kindPtr
	options.Verbose = *verbosePtr
	options.Year = *yearPtr
	options.Title = *titlePtr
	parsed = true
}