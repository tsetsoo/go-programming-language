package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type multiTierSort struct {
	t               []*Track
	fieldsToCompare []string
}

func (s multiTierSort) Len() int      { return len(s.t) }
func (s multiTierSort) Swap(i, j int) { s.t[i], s.t[j] = s.t[j], s.t[i] }
func (s multiTierSort) Less(i, j int) bool {
	for fieldIndex := 0; fieldIndex < len(s.fieldsToCompare); fieldIndex++ {
		switch s.fieldsToCompare[fieldIndex] {
		case "Title":
			if s.t[i].Title != s.t[j].Title {
				return s.t[i].Title < s.t[j].Title
			} else {
				continue
			}
		case "Artist":
			if s.t[i].Artist != s.t[j].Artist {
				return s.t[i].Artist < s.t[j].Artist
			} else {
				continue
			}
		case "Album":
			if s.t[i].Album != s.t[j].Album {
				return s.t[i].Album < s.t[j].Album
			} else {
				continue
			}
		case "Year":
			if s.t[i].Year != s.t[j].Year {
				return s.t[i].Year < s.t[j].Year
			} else {
				continue
			}
		case "Length":
			if s.t[i].Length != s.t[j].Length {
				return s.t[i].Length < s.t[j].Length
			} else {
				continue
			}
		}
	}
	return false
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	printTracks(tracks)
	fmt.Println()
	sort.Sort(multiTierSort{tracks, []string{"Title", "Year", "Length"}})
	printTracks(tracks)
}
