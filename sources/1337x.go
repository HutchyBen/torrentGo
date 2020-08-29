package sources

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
)

func LEET(c *cli.Context) error {
	searchTerm := c.Args().Get(0)

	url := "https://1337x.to/search/" + searchTerm + "/1/"
	var selections []Torrent
	doc := DocumentFromURL(url)

	doc.Find("tbody").First().Children().Each(func(i int, selection *goquery.Selection) {
		torrent := Torrent{}
		torrentName := selection.Find(".name").Last().Text()
		torrent.name = torrentName
		torrent.author = selection.Find(".coll-5").First().Children().First().Text()
		torrent.seeders, _ = strconv.Atoi(selection.Find(".coll-2").First().Text())
		torrent.leechers, _ = strconv.Atoi(selection.Find(".coll-3").First().Text())
		size := selection.Find(".coll-4").First().Text()
		size = size[:strings.Index(size, "B")+1]
		size = strings.Replace(size, " ", "", -1)
		torrent.size = size
		href, _ := selection.Find(".coll-1").Children().Eq(1).Attr("href")
		torrent.pageURL = "https://1337x.to" + href

		selections = append(selections, torrent)
	})
	if len(selections) == 0 {
		return cli.Exit("There is no found torrents", 69)
	}

	index := DisplayMenu(selections)
	fmt.Print("\n")
	selectedTorrent := selections[index]
	response := ""
	survey.AskOne(&survey.Select{Message: "What do you want to do?", Options: []string{"Download", "Goto Torrent Page", "Cancel"}}, &response)

	switch response {
	case "Cancel":
		cli.Exit("User cancelled", 0)
		break

	case "Goto Torrent Page":
		browser.OpenURL(selectedTorrent.pageURL)
		break

	case "Download":
		magnet := leetGetTorrent(selectedTorrent.pageURL)

		DownloadFile(magnet)
	}

	return nil
}

func leetGetTorrent(url string) string {
	doc := DocumentFromURL(url)
	fmt.Println()
	data, _ := doc.Find(".torrent-detail-page").First().Children().Eq(1).Children().First().Children().First().Children().First().Children().First().Attr("href")
	return data
}
