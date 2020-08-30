package sources

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/PuerkitoBio/goquery"
	"github.com/anacrolix/torrent"
	"log"
	"net/http"
	"strconv"
)

func DocumentFromURL(url string) *goquery.Document {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetDisplayName(torrent Torrent) string {
	data := ""
	if len(torrent.name) > 70 {
		data += torrent.name[:70]
	} else {
		data += torrent.name
	}

	data += " By " + torrent.author
	data += " Size: " + torrent.size
	data += " S: " + strconv.Itoa(torrent.seeders)
	data += " L: " + strconv.Itoa(torrent.leechers)
	return data
}

func DisplayMenu(baseURL string, page *int, forw bool, items []Torrent) int {
	var displayed []string

	if *page > 1 {
		displayed = append(displayed, ">>>BACK<<<")
	}

	for _, v := range items {
		displayed = append(displayed, GetDisplayName(v))
	}

	if forw {
		displayed = append(displayed, ">>>FORWARD<<<")

	}

	prompt := &survey.Select{
		Message: "Choose a torrent:",
		Options: displayed,
	}
	var choice string

	survey.AskOne(prompt, &choice)
	fmt.Println(choice == ">>>FORWARD<<<")
	if choice == ">>>FORWARD<<<" {
		*page++
		return 1000
	} else if choice == ">>>BACK<<<" {
		*page--
		return 999
	}
	var index int
	for i, v := range items {
		if GetDisplayName(v) == choice {
			index = i
			break
		}
	}

	return index
}

func DownloadFile(url string) {
	fmt.Println("It may look like the download is stuck")
	c, _ := torrent.NewClient(nil)
	defer c.Close()
	t, _ := c.AddMagnet(url)
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
}
