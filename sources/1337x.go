package sources

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
	"log"
)

func LEET(c *cli.Context) error{
	searchTerm := c.Args().Get(0)

	url := "https://1337x.to/search/" + searchTerm +"/1/"
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

	index, err := DisplayMenu(selections)
	if err != nil {
		log.Fatal(err)
	}



	attr, _ := doc.Find("tbody").First().Children().Eq(index).Children().First().Children().Eq(1).Attr("href")
	magnet := leetGetTorrent("https://1337x.to" + attr)

	DownloadFile(magnet)
	fmt.Println(magnet)
	return nil
}

func leetGetTorrent(url string) string{
	fmt.Println(url)
	doc := DocumentFromURL(url)
	fmt.Println()
	data, _ := doc.Find(".l6bc4e773c6fde2c9931634044cb0c8f8aa1a32f6").First().Attr("href")
	return data
}
