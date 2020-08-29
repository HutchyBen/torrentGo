package sources

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

func LEET(c *cli.Context) error {
	searchTerm := c.Args().Get(0)

	url := "https://1337x.to/search/" + searchTerm + "/1/"
	var selections []string
	doc := DocumentFromURL(url)

	doc.Find("tbody").First().Children().Each(func(i int, selection *goquery.Selection) {
		torrentName := selection.Find(".name").Last().Text()

		if len(torrentName) > 85 {
			torrentName = torrentName[:85] + "..."
		}

		torrentAuthor := selection.Find(".coll-5").First().Children().First().Text()

		torrentName = torrentName + " By " + torrentAuthor
		selections = append(selections, torrentName)
	})
	if len(selections) == 0 {
		return cli.Exit("There is no found torrents", 69)
	}

	index, err := DisplayMenu(selections)
	if err != nil {
		return cli.Exit("Could not display menu", 69)
	}

	attr, _ := doc.Find("tbody").First().Children().Eq(index).Children().First().Children().Eq(1).Attr("href")
	magnet := leetGetTorrent("https://1337x.to" + attr)

	DownloadFile(magnet)
	fmt.Println(magnet)
	return nil
}

func leetGetTorrent(url string) string {
	doc := DocumentFromURL(url)
	fmt.Println()
	data, _ := doc.Find(".torrent-detail-page").First().Children().Eq(1).Children().First().Children().First().Children().First().Children().First().Attr("href")
	return data
}
