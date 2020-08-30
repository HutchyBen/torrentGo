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

var page int

func LEET(c *cli.Context) error {
	page = 1
	searchTerm := c.Args().Get(0)

	category := c.String("category")
	sorting := c.String("sort")
	url, err := leetMakeURL(searchTerm, sorting, category)
	if err != "" {
		return cli.Exit(err, 69)
	}
	out := leetMain(url, 1)
	if out != nil {
		return out
	}

	return cli.Exit("Program finished successfully  (probably)", 0)
}

func leetMain(baseurl string, page int) cli.ExitCoder {

	var selections []Torrent

	doc := DocumentFromURL(baseurl + strconv.Itoa(page) + "/")
	// Fetch all torrents on page
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

	// Check if there is other pages to be navigated
	foward := doc.Find(".last").Size() > 0

	index := DisplayMenu(baseurl, &page, foward, selections)
	fmt.Print("\n")

	if index == 999 {
		return leetMain(baseurl, page)
	} else if index == 1000 {
		return leetMain(baseurl, page)
	}

	selectedTorrent := selections[index]
	response := ""
	survey.AskOne(&survey.Select{Message: "What do you want to do?", Options: []string{"Download", "Goto Torrent Page", "Back", "Cancel"}}, &response)

	switch response {
	case "Cancel":
		return cli.Exit("User cancelled", 0)
		break

	case "Goto Torrent Page":
		browser.OpenURL(selectedTorrent.pageURL)
		break

	case "Download":
		magnet := leetGetTorrent(selectedTorrent.pageURL)

		DownloadFile(magnet)

	case "Back":
		return leetMain(baseurl, page)
	}

	return nil
}

func leetMakeURL(search string, sorting string, category string) (string, string) {
	url := "https://1337x.to/"

	if sorting != "" {
		url += "sort-"
	}

	if category != "" {
		url += "category-"
	}

	url += "search/" + search
	category = strings.ToLower(category)
	isValid := false
	for _, v := range []string{"tv", "movies", "games", "music", "applications", "anime", "xxx", "other", "documentaries", ""} {
		if category == v {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", "Not a correct category"
	}

	if category == "" {
		// ok so its 2 am and i dont know if this is smart
	} else if category == "tv" || category == "xxx" {
		url += "/" + strings.ToUpper(category)
	} else {
		category = strings.ToUpper(string(category[0])) + category[1:]
		url += "/" + category
	}

	isValid = false
	sorting = strings.ToLower(sorting)
	for _, v := range []string{"sizea", "sized", "timea", "timed", "seedersa", "seedersd", "leechersa", "leechersd", ""} {
		if sorting == v {
			isValid = true
			break
		}

	}
	if !isValid {
		return "", "Not a valid sorting"
	}

	if sorting == "" {
		url += "/"
	} else {
		sortingName := sorting[:len(sorting)-1]

		url += "/" + sortingName

		order := sorting[len(sorting)-1:]

		if order == "a" {
			url += "/asc/"
		} else {
			url += "/desc/"
		}
	}

	return url, ""
}

func leetGetTorrent(url string) string {
	doc := DocumentFromURL(url)
	fmt.Println()
	data, _ := doc.Find(".torrent-detail-page").First().Children().Eq(1).Children().First().Children().First().Children().First().Children().First().Attr("href")
	return data
}
